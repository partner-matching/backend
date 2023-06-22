package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
)

type PartnerRepo interface {
	AddTeam(ctx context.Context, team *CreateTeam) error
	DeleteTeam(ctx context.Context, teamId int32) error
	UpdateTeam(ctx context.Context, team *UpdateTeam) error
	GetTeam(ctx context.Context, teamId int32) (*Team, error)
	GetTeamList(ctx context.Context, query *TeamQuery, page, pageSizer int32) ([]*Team, error)
}

type PartnerRepoUseCase struct {
	repo PartnerRepo
	log  *log.Helper
	re   Recovery
	tm   Transaction
}

type Team struct {
	Id          int32  `validate:"omitempty,gte=0" comment:"id"`
	Name        string `validate:"omitempty,min=1,max=256" comment:"队伍名"`
	Description string `validate:"omitempty,max=1024" comment:"描述"`
	MaxNum      int32  `validate:"omitempty,gte=1" comment:"最大人数"`
	ExpireTime  string `validate:"omitempty,datetime" comment:"过期时间"`
	UserId      int32  `validate:"omitempty,gte=0" comment:"用户id"`
	Status      int32  `validate:"omitempty,oneof=0 1 2" comment:"状态"`
	Password    string `validate:"omitempty,max=512" comment:"密码"`
}

func NewPartnerRepoUseCase(repo PartnerRepo, re Recovery, tm Transaction, logger log.Logger) *PartnerRepoUseCase {
	return &PartnerRepoUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "user/biz/PartnerRepoUseCase")),
		tm:   tm,
		re:   re,
	}
}

// CreateTeam 校验信息
//1. 队伍人数 > 1 且 <= 20
//2. 队伍标题 <= 20
//3. 描述 <= 512
//4. status 是否公开（int）不传默认为 0（公开）
//5. 如果 status 是加密状态，一定要有密码，且密码 <= 32
//6. 超时时间 > 当前时间
//7. 校验用户最多创建 5 个队伍
type CreateTeam struct {
	Team
	Name   string `validate:"required,min=1,max=256" comment:"队伍名"`
	MaxNum int32  `validate:"required,gte=1" comment:"最大人数"`
	Status int32  `validate:"required,oneof=0 1 2" comment:"状态"`
}

// AddTeam 是否登录，未登录不允许创建
func (r *PartnerRepoUseCase) AddTeam(ctx context.Context, team *CreateTeam) error {
	err := r.repo.AddTeam(ctx, team)
	if err != nil {
		return v1.ErrorAddTeamFailed("%s", err.Error())
	}
	return nil
}

type DeleteTeam struct {
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) DeleteTeam(ctx context.Context, team *DeleteTeam) error {
	err := r.repo.DeleteTeam(ctx, team.Id)
	if err != nil {
		return v1.ErrorDeleteTeamFailed("%s", err.Error())
	}
	return nil
}

type UpdateTeam struct {
	Team
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) UpdateTeam(ctx context.Context, team *UpdateTeam) error {
	err := r.repo.UpdateTeam(ctx, team)
	if err != nil {
		return v1.ErrorUpdateTeamFailed("%s", err.Error())
	}
	return nil
}

type GetTeam struct {
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) GetTeam(ctx context.Context, team *GetTeam) (*Team, error) {
	t, err := r.repo.GetTeam(ctx, team.Id)
	if err != nil {
		return nil, v1.ErrorGetTeamFailed("%s", err.Error())
	}
	return t, nil
}

type TeamQuery struct {
	Name       string
	ExpireTime string
	Status     int32
}

func (r *PartnerRepoUseCase) GetTeamList(ctx context.Context, team *TeamQuery, page, pageSize int32) ([]*Team, error) {
	list, err := r.repo.GetTeamList(ctx, team, page, pageSize)
	if err != nil {
		return nil, v1.ErrorGetTeamListFailed("%s", err.Error())
	}
	return list, nil
}
