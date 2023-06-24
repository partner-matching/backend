package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewUserService, NewPartnerService)

type PartnerService struct {
	v1.UnimplementedPartnerServiceServer
	uc  *biz.UserUseCase
	ac  *biz.AuthRepoUseCase
	pc  *biz.PartnerRepoUseCase
	vc  *biz.ValidateUseCase
	log *log.Helper
}

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc  *biz.UserUseCase
	ac  *biz.AuthRepoUseCase
	pc  *biz.PartnerRepoUseCase
	vc  *biz.ValidateUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, ac *biz.AuthRepoUseCase, pc *biz.PartnerRepoUseCase, vc *biz.ValidateUseCase, logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper(log.With(logger, "module", "user/service")),
		uc:  uc,
		ac:  ac,
		vc:  vc,
		pc:  pc,
	}
}

func NewPartnerService(uc *biz.UserUseCase, ac *biz.AuthRepoUseCase, pc *biz.PartnerRepoUseCase, vc *biz.ValidateUseCase, logger log.Logger) *PartnerService {
	return &PartnerService{
		log: log.NewHelper(log.With(logger, "module", "user/partner")),
		uc:  uc,
		ac:  ac,
		vc:  vc,
		pc:  pc,
	}
}
