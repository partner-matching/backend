package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

var ProviderSet = wire.NewSet(NewUserService)

type PartnerService struct {
	v1.UnimplementedUserServiceServer
	uc  *biz.UserUseCase
	ac  *biz.AuthRepoUseCase
	pc  *biz.PartnerRepoUseCase
	vc  *biz.ValidateUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, ac *biz.AuthRepoUseCase, pc *biz.PartnerRepoUseCase, vc *biz.ValidateUseCase, logger log.Logger) *PartnerService {
	return &PartnerService{
		log: log.NewHelper(log.With(logger, "module", "user/service")),
		uc:  uc,
		ac:  ac,
		vc:  vc,
		pc:  pc,
	}
}

func (s *PartnerService) GetHealth(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
