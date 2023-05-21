package data

import (
	"context"
	"fmt"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"github.com/pkg/errors"
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
		return errors.Wrapf(err, fmt.Sprintf("fail to delete user: userId(%s)", userId))
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
