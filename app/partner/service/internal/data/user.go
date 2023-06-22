package data

import (
	"context"
	"encoding/json"
	"fmt"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"github.com/pkg/errors"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user/data/user")),
	}
}

func (r *userRepo) GetUsersList(ctx context.Context, pageNum, pageSize int32) ([]*biz.User, error) {
	if pageNum == 0 {
		pageNum = 1
	}
	usersList, err := r.getUserListFromCache(ctx, pageNum, pageSize)
	if err == nil {
		return usersList, nil
	}
	if !errors.Is(err, redis.Nil) {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get users list from cache"))
	}

	usersList, err = r.getUserListFromDb(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	if len(usersList) > 0 {
		r.setUserListToCache(ctx, usersList, pageNum, pageSize)
	}

	return usersList, nil
}

func (r *userRepo) getUserListFromCache(ctx context.Context, pageNum, pageSize int32) ([]*biz.User, error) {
	key := fmt.Sprintf("partner:user:recommend:%v:%v", pageNum, pageSize)
	result, err := r.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	marshalList := make([]string, 0)
	err = json.Unmarshal([]byte(result), &marshalList)
	if err != nil {
		return nil, err
	}

	usersList := make([]*biz.User, 0)
	for _, item := range marshalList {
		user := &biz.User{}
		err = user.UnmarshalJSON([]byte(item))
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (r *userRepo) getUserListFromDb(ctx context.Context, pageNum, pageSize int32) ([]*biz.User, error) {
	list := make([]*User, 0)
	err := r.data.db.WithContext(ctx).Where("id >= ? and isDelete = 0", (pageNum-1)*pageSize).Limit(int(pageSize)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get users list"))
	}

	var usersList []*biz.User
	for _, item := range list {
		user := &biz.User{}
		util.StructAssign(user, item)
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (r *userRepo) setUserListToCache(ctx context.Context, usersList []*biz.User, pageNum, pageSize int32) {
	var marshalList []string
	for _, item := range usersList {
		marshal, err := item.MarshalJSON()
		if err != nil {
			r.log.Errorf("fail to set user info to json: json.Marshal(%v), error(%v)", item, err)
			continue
		}
		marshalList = append(marshalList, string(marshal))
	}

	key := fmt.Sprintf("partner:user:recommend:%v:%v", pageNum, pageSize)
	marshal, err := json.Marshal(marshalList)
	if err != nil {
		r.log.Errorf("fail to set user list to json: json.Marshal(%v), error(%v)", usersList, err)
	}
	err = r.data.redisCli.SetNX(ctx, key, string(marshal), time.Minute*1).Err()
	if err != nil {
		r.log.Errorf("fail to set user list to cache: redis.Set(%v), error(%v)", usersList, err)
	}
}

func (r *userRepo) GetUserRole(ctx context.Context) (int32, int32, error) {
	user, err := r.getUserFromSession(ctx)
	if err != nil {
		return 0, 0, err
	}
	return user.Id, user.Role, nil
}

func (r *userRepo) GetUserSession(ctx context.Context) (*biz.User, error) {
	user, err := r.getUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	result := &biz.User{}
	util.StructAssign(result, user)
	return result, nil
}

func (r *userRepo) GetCurrentUser(ctx context.Context, userId int32) (*biz.User, error) {
	user := &User{
		Id: userId,
	}
	err := r.data.db.WithContext(ctx).Where("id = ?", userId).First(user).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("get user failed: userId(%v)", userId))
	}

	result := &biz.User{}
	util.StructAssign(result, user)
	return result, nil
}

// SearchUsers 查询用户（允许根据用户名查询，仅管理员可查询）
func (r *userRepo) SearchUsers(ctx context.Context, userName string) ([]*biz.User, error) {
	list := make([]*User, 0)
	var err error
	switch userName {
	case "":
		err = r.data.db.WithContext(ctx).Where("isDelete = 0").Find(&list).Error
	default:
		err = r.data.db.WithContext(ctx).Where("userName like ? and isDelete = 0", userName).Find(&list).Error
	}
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to search users: userName(%s)", userName))
	}

	var search []*biz.User
	for _, item := range list {
		user := &biz.User{}
		util.StructAssign(user, item)
		search = append(search, user)
	}
	return search, nil
}

// SearchUsersByTags 根据标签搜索用户
func (r *userRepo) SearchUsersByTags(ctx context.Context, tagList []string) ([]*biz.User, error) {
	list := make([]*User, 0)
	queryDB := r.data.db.WithContext(ctx).Where("isDelete = 0")
	// 链式查询
	for _, tag := range tagList {
		queryDB = queryDB.Where("tags like ?", "%"+tag+"%")
	}
	err := queryDB.Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to search users by tags: tagList(%v)", tagList))
	}

	var search []*biz.User
	for _, item := range list {
		user := &biz.User{}
		util.StructAssign(user, item)
		search = append(search, user)
	}
	return search, nil
}

// UpdateUser 更新用户信息
func (r *userRepo) UpdateUser(ctx context.Context, update *biz.UpdateUser) error {
	user := &User{
		UserName:  update.UserName,
		AvatarUrl: update.AvatarUrl,
		Gender:    update.Gender,
		Phone:     update.Phone,
		Email:     update.Email,
	}
	err := r.data.db.WithContext(ctx).Where("id = ?", update.Id).Updates(user).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to update user: user(%v)", update))
	}
	return nil
}

// DeleteUser 删除用户
func (r *userRepo) DeleteUser(ctx context.Context, userId int32) error {
	user := &User{}
	user.Id = userId
	user.IsDelete = 1
	err := r.data.db.WithContext(ctx).Where("id = ? and isDelete = 0", userId).Delete(user).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to delete user: userId(%v)", userId))
	}
	return nil
}

func (r *userRepo) getUserFromSession(ctx context.Context) (*User, error) {
	req, _ := util.GetRequestAndResponse(ctx)
	session, err := r.data.sessionStore.Get(req, r.data.conf.UserLoginState)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("get session failed"))
	}
	if len(session.Values) == 0 {
		return nil, kerrors.NotFound("user not found from session", "")
	}
	user := string(session.Values["user"].([]uint8))
	var storeUser = &User{}
	err = storeUser.UnmarshalJSON([]byte(user))
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("json unmarshal error: user(%v)", user))
	}
	return storeUser, nil
}
