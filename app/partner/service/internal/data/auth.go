package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user/data/register")),
	}
}

func (r *authRepo) AccountExist(ctx context.Context, userAccount string) (bool, error) {
	user := &User{}
	err := r.data.db.WithContext(ctx).Select("id").Where("userAccount = ?", userAccount).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrapf(err, fmt.Sprintf("fail to get account: userAccount(%v)", userAccount))
	}
	return true, nil
}

func (r *authRepo) UserRegister(ctx context.Context, userAccount, passwordHash string) (int32, error) {
	user := &User{
		UserAccount:  userAccount,
		UserPassword: passwordHash,
	}
	err := r.data.db.WithContext(ctx).Select("userAccount", "userPassword").Create(user).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to register user: userAccount(%s), userPassword(%s)", userAccount, passwordHash))
	}
	return user.Id, nil
}

func (r *authRepo) UserLogin(ctx context.Context, userAccount, passwordHash string) (*biz.User, error) {
	user := &User{
		UserAccount:  userAccount,
		UserPassword: passwordHash,
	}
	err := r.data.db.WithContext(ctx).Where("userAccount = ? and userPassword = ? and isDelete = 0", userAccount, passwordHash).First(user).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("user login failed: userAccount(%s), userPassword(%s)", userAccount, passwordHash))
	}

	result := &biz.User{}
	util.StructAssign(result, user)
	return result, nil
}

func (r *authRepo) UserLogout(ctx context.Context) error {
	req, res := util.GetRequestAndResponse(ctx)
	session, err := r.data.sessionStore.Get(req, r.data.conf.UserLoginState)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("get session failed"))
	}
	session.Options.MaxAge = -1
	err = session.Save(req, res)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("set session failed"))
	}
	return nil
}

func (r *authRepo) SetLoginSession(ctx context.Context, user *biz.User) error {
	marshal, err := user.MarshalJSON()
	if err != nil {
		r.log.Errorf("fail to set user info to json: json.Marshal(%v), error(%v)", user, err)
		return nil
	}
	req, res := util.GetRequestAndResponse(ctx)
	session, err := r.data.sessionStore.Get(req, r.data.conf.UserLoginState)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("get session failed: user(%v)", user))
	}
	session.Values["user"] = marshal
	err = session.Save(req, res)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("set session failed: user(%v)", user))
	}
	return nil
}
