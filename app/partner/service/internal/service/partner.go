package service

import (
	"context"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PartnerService) AddTeam(ctx context.Context, req *v1.Team) (*emptypb.Empty, error) {
	//是否登录，未登录不允许创建
	user, err := s.uc.IsUserLogin(ctx)
	if err != nil {
		return nil, err
	}

	team := &biz.CreateTeam{}
	util.StructAssign(team, req)

	team.UserId = user.Id
	err = s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}

	err = s.pc.AddTeam(ctx, user, team)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) DeleteTeam(ctx context.Context, req *v1.DeleteTeamReq) (*emptypb.Empty, error) {
	//是否登录，未登录不允许创建
	_, err := s.uc.IsUserLogin(ctx)
	if err != nil {
		return nil, err
	}

	team := &biz.DeleteTeam{
		Id: req.Id,
	}
	err = s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}

	err = s.pc.DeleteTeam(ctx, team)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) UpdateTeam(ctx context.Context, req *v1.UpdateTeamReq) (*emptypb.Empty, error) {
	//是否登录，未登录不允许修改
	user, err := s.uc.IsUserLogin(ctx)
	if err != nil {
		return nil, err
	}

	//是否是管理员
	isAdmin := true
	err = s.uc.IsAdmin(ctx)
	if err != nil {
		isAdmin = false
	}

	team := &biz.UpdateTeam{}
	util.StructAssign(team, req)

	err = s.vc.ParamsValidate(team)
	if err != nil {
		return nil, err
	}

	err = s.pc.UpdateTeam(ctx, team, user.Id, isAdmin)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) GetTeam(ctx context.Context, req *v1.GetTeamReq) (*v1.GetTeamResponse, error) {
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
	//只有管理员才能查看加密还有非公开的房间
	if req.Query.Status != 0 {
		err := s.uc.IsAdmin(ctx)
		if err != nil {
			return nil, err
		}
	}
	team := &biz.TeamQuery{}
	util.StructAssign(team, req.Query)

	teamList, err := s.pc.GetTeamList(ctx, team, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetTeamListResponse{}
	for _, item := range teamList {
		reply.Data = append(reply.Data, &v1.GetTeamListResponse_TeamInfo{
			Team: &v1.Team{
				Id:          item.Id,
				Name:        item.Name,
				MaxNum:      item.MaxNum,
				ExpireTime:  item.ExpireTime,
				UserId:      item.UserId,
				Status:      item.Status,
				Description: item.Description,
			},
			UserInfo: &v1.User{
				Id:          item.UserInfo.Id,
				UserName:    item.UserInfo.UserName,
				UserAccount: item.UserInfo.UserAccount,
				AvatarUrl:   item.UserInfo.AvatarUrl,
				Phone:       item.UserInfo.Phone,
				Email:       item.UserInfo.Email,
				UserStatus:  item.UserInfo.UserStatus,
				Gender:      item.UserInfo.Gender,
				UserRole:    item.UserInfo.Role,
				CreateTime:  item.UserInfo.CreateTime.String(),
				Tags:        item.UserInfo.Tags,
				Profile:     item.UserInfo.Profile,
			},
		})
	}
	return reply, nil
}

func (s *PartnerService) JoinTeam(ctx context.Context, req *v1.JoinTeamReq) (*emptypb.Empty, error) {
	//是否登录，未登录不允许加入队伍
	user, err := s.uc.IsUserLogin(ctx)
	if err != nil {
		return nil, err
	}

	joinTeam := &biz.JoinTeam{}
	util.StructAssign(joinTeam, req)
	err = s.vc.ParamsValidate(joinTeam)
	if err != nil {
		return nil, err
	}

	err = s.pc.JoinTeam(ctx, joinTeam, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PartnerService) QuitTeam(ctx context.Context, req *v1.QuitTeamReq) (*emptypb.Empty, error) {
	//是否登录，未登录不允许修改
	user, err := s.uc.IsUserLogin(ctx)
	if err != nil {
		return nil, err
	}

	quitTeam := &biz.QuitTeam{}
	util.StructAssign(quitTeam, req)
	err = s.vc.ParamsValidate(quitTeam)
	if err != nil {
		return nil, err
	}

	err = s.pc.QuitTeam(ctx, quitTeam, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
