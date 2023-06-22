package service

import (
	"context"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PartnerService) AddTeam(ctx context.Context, req *v1.Team) (*emptypb.Empty, error) {
	team := &biz.CreateTeam{}
	util.StructAssign(team, req)
	err := s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}
	err = s.pc.AddTeam(ctx, team)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) DeleteTeam(ctx context.Context, req *v1.Team) (*emptypb.Empty, error) {
	team := &biz.DeleteTeam{
		Id: req.Id,
	}
	err := s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}
	err = s.pc.DeleteTeam(ctx, team)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) UpdateTeam(ctx context.Context, req *v1.Team) (*emptypb.Empty, error) {
	team := &biz.UpdateTeam{
		Id: req.Id,
	}
	err := s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}
	err = s.pc.UpdateTeam(ctx, team)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) GetTeam(ctx context.Context, req *v1.Team) (*v1.GetTeamResponse, error) {
	team := &biz.GetTeam{
		Id: req.Id,
	}
	err := s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}
	result, err := s.pc.GetTeam(ctx, team)
	if err != nil {
		return nil, err
	}

	return &v1.GetTeamResponse{
		Data: &v1.Team{
			Name:        result.Name,
			MaxNum:      result.MaxNum,
			ExpireTime:  result.ExpireTime,
			UserId:      result.UserId,
			Status:      result.Status,
			Password:    result.Password,
			Description: result.Description,
		},
	}, nil
}

func (s *PartnerService) GetTeamList(ctx context.Context, req *v1.GetTeamListReq) (*v1.GetTeamListResponse, error) {
	team := &biz.TeamQuery{}
	util.StructAssign(team, req.Query)
	teamList, err := s.pc.GetTeamList(ctx, team, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetTeamListResponse{}
	for _, item := range teamList {
		reply.Data = append(reply.Data, &v1.Team{
			Id:          item.Id,
			Name:        item.Name,
			MaxNum:      item.MaxNum,
			ExpireTime:  item.ExpireTime,
			UserId:      item.UserId,
			Status:      item.Status,
			Password:    item.Password,
			Description: item.Description,
		})
	}
	return reply, nil
}
