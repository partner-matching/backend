package biz

import (
	"context"
	"encoding/json"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/conf"
	"github.com/pkg/errors"
	"time"
	//v1 "github.com/partner-matching/backend/api/user/service/v1"
)

type UserRepo interface {
	SearchUsers(ctx context.Context, userName string) ([]*User, error)
	SearchUsersByTags(ctx context.Context, tagList []string) ([]*User, error)
	UpdateUser(ctx context.Context, update *UpdateUser) error
	DeleteUser(ctx context.Context, userName int32) error
	GetUserRole(ctx context.Context) (int32, int32, error)
	GetUserSession(ctx context.Context) (*User, error)
	GetCurrentUser(ctx context.Context, userId int32) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
	re   Recovery
	tm   Transaction
	conf *conf.UserConstant
}

//easyjson:json
type User struct {
	Id           int32
	UserName     string
	UserAccount  string
	AvatarUrl    string
	Gender       int32
	UserPassword string
	Phone        string
	Email        string
	UserStatus   int32
	Role         int32
	CreateTime   time.Time
	Tags         string
	Profile      string
}

type SearchUser struct {
	UserName string
}

type DeleteUser struct {
	Id int32 `validate:"required,gt=0" comment:"用户Id"`
}

type UpdateUser struct {
	Id        int32
	UserName  string `validate:"omitempty" comment:"用户昵称"`
	AvatarUrl string `validate:"omitempty" comment:"用户头像"`
	Gender    int32  `validate:"omitempty" comment:"性别"`
	Phone     string `validate:"omitempty" comment:"手机号"`
	Email     string `validate:"omitempty" comment:"邮箱"`
}

func NewUserUseCase(repo UserRepo, re Recovery, tm Transaction, logger log.Logger, conf *conf.UserConstant) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "user/biz/userUseCase")),
		tm:   tm,
		re:   re,
		conf: conf,
	}
}

func (i *User) DoValidate(trans ut.Translator, validate *validator.Validate) error {
	err := validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}

// SearchUsers 用户搜索逻辑
//1. 判断是否有管理员权限
//2. 根据用户名进行模糊查询
func (r *UserUseCase) SearchUsers(ctx context.Context, userName string) ([]*User, error) {
	err := r.isAdmin(ctx)
	if err != nil {
		return nil, err
	}

	users, err := r.repo.SearchUsers(ctx, userName)
	if err != nil {
		return nil, v1.ErrorUserSearchFailed("%s", err.Error())
	}
	return users, nil
}

// SearchUsersByTags 根据标签搜索用户
func (r *UserUseCase) SearchUsersByTags(ctx context.Context, tagList []string) ([]*User, error) {
	if len(tagList) == 0 {
		return nil, v1.ErrorUserTagsEmpty("user search tags is empty")
	}
	users, err := r.repo.SearchUsersByTags(ctx, tagList)
	if err != nil {
		return nil, v1.ErrorUserSearchFailed("%s", err.Error())
	}
	return users, nil
}

// SearchUsersByTagsInCache 根据标签搜索用户（在内存中判断）
func (r *UserUseCase) SearchUsersByTagsInCache(ctx context.Context, tagList []string) ([]*User, error) {
	// 先把所有用户从数据库中取出
	users, err := r.repo.SearchUsers(ctx, "")
	if err != nil {
		return nil, v1.ErrorUserSearchFailed("%s", err.Error())
	}

	result := make([]*User, 0)
	// 将tags字段取出，解析json字符串
OuterLoop:
	for _, user := range users {
		// tags字段为空，则直接下一个
		if user.Tags == "" {
			continue
		}
		tagsMap := make(map[string]int, 0)
		err = json.Unmarshal([]byte(user.Tags), &tagsMap)
		if err != nil {
			return nil, v1.ErrorUnknownError("%s", err.Error())
		}
		for _, tag := range tagList {
			// 若有一个标签不存在，则下一个
			if _, ok := tagsMap[tag]; !ok {
				continue OuterLoop
			}
		}
		// 标签全符合，才是最后搜索结果
		result = append(result, user)
	}
	return result, nil
}

// DeleteUser 用户删除逻辑
//1. 判断是否有管理员权限
//2. 根据用户id进行逻辑删除
func (r *UserUseCase) DeleteUser(ctx context.Context, userId int32) error {
	err := r.isAdmin(ctx)
	if err != nil {
		return err
	}

	err = r.repo.DeleteUser(ctx, userId)
	if err != nil {
		return v1.ErrorUserDeleteFailed("%s", err.Error())
	}
	return nil
}

// GetCurrentUser 当前登录用户获取逻辑
//1. 判断session是否存在
//2. 如果存在，从数据库中获取最新用户信息返回
func (r *UserUseCase) GetCurrentUser(ctx context.Context) (*User, bool, error) {
	user, exist, err := r.isSessionExist(ctx)
	if err != nil {
		return nil, false, err
	}
	if !exist {
		return nil, true, err
	}
	user, err = r.repo.GetCurrentUser(ctx, user.Id)
	if err != nil {
		return nil, false, v1.ErrorUnknownError("%s", err.Error())
	}
	return user, false, nil
}

// UpdateUser 更新用户信息
func (r *UserUseCase) UpdateUser(ctx context.Context, updateUser *UpdateUser) error {
	// 没登录或获取session出错，直接退出
	user, exist, err := r.isSessionExist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return v1.ErrorLoginStateTimeout("")
	}

	// 既不是管理员，也不是修改自己的，直接退出
	if err = r.isAdmin(ctx); err != nil && user.Id != updateUser.Id {
		return v1.ErrorPermissionDeny("update user: %v", updateUser)
	}

	updateUser.Id = user.Id
	err = r.repo.UpdateUser(ctx, updateUser)
	if err != nil {
		return v1.ErrorUpdateUser("%s", err.Error())
	}
	return nil
}

func (r *UserUseCase) isAdmin(ctx context.Context) error {
	userId, role, err := r.repo.GetUserRole(ctx)
	if kerrors.IsNotFound(err) {
		return v1.ErrorLoginStateTimeout("")
	}

	if err != nil {
		return v1.ErrorUnknownError("%s", err.Error())
	}

	if role != r.conf.AdminRole {
		return v1.ErrorPermissionDeny("userId: %s, userRole: %v", userId, role)
	}
	return nil
}

func (r *UserUseCase) isSessionExist(ctx context.Context) (*User, bool, error) {
	user, err := r.repo.GetUserSession(ctx)
	if kerrors.IsNotFound(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, v1.ErrorUnknownError("%s", err.Error())
	}
	return user, true, nil
}