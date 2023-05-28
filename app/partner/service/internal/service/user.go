package service

import (
	"context"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PartnerService) SearchUsers(ctx context.Context, req *v1.SearchUsersReq) (*v1.SearchUsersReply, error) {
	usersList, err := s.uc.SearchUsers(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	reply := &v1.SearchUsersReply{
		Data: make([]*v1.User, 0, len(usersList)),
	}
	for _, item := range usersList {
		reply.Data = append(reply.Data, &v1.User{
			Id:          item.Id,
			UserName:    item.UserName,
			UserAccount: item.UserAccount,
			AvatarUrl:   item.AvatarUrl,
			Phone:       item.Phone,
			Email:       item.Email,
			UserStatus:  item.UserStatus,
			Gender:      item.Gender,
			UserRole:    item.Role,
			CreateTime:  item.CreateTime.String(),
			Tags:        item.Tags,
			Profile:     item.Profile,
		})
	}
	return reply, nil
}

// SearchUsersByTags 根据标签搜索用户
func (s *PartnerService) SearchUsersByTags(ctx context.Context, req *v1.SearchUsersByTagsReq) (*v1.SearchUsersReply, error) {
	usersList, err := s.uc.SearchUsersByTagsInCache(ctx, req.TagNameList)
	if err != nil {
		return nil, err
	}

	reply := &v1.SearchUsersReply{
		Data: make([]*v1.User, 0, len(usersList)),
	}
	for _, item := range usersList {
		reply.Data = append(reply.Data, &v1.User{
			Id:          item.Id,
			UserName:    item.UserName,
			UserAccount: item.UserAccount,
			AvatarUrl:   item.AvatarUrl,
			Phone:       item.Phone,
			Email:       item.Email,
			UserStatus:  item.UserStatus,
			Gender:      item.Gender,
			UserRole:    item.Role,
			CreateTime:  item.CreateTime.String(),
			Tags:        item.Tags,
			Profile:     item.Profile,
		})
	}
	return reply, nil
}

func (s *PartnerService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*emptypb.Empty, error) {
	search := &biz.DeleteUser{
		Id: req.Id,
	}
	err := s.vc.ParamsValidate(search)
	if err != nil {
		return nil, err
	}

	err = s.uc.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) GetCurrentUser(ctx context.Context, _ *emptypb.Empty) (*v1.GetCurrentReply, error) {
	user, empty, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	if empty {
		return &v1.GetCurrentReply{
			Data: &v1.User{
				Empty: empty,
			},
		}, nil
	}
	return &v1.GetCurrentReply{
		Data: &v1.User{
			Empty:       false,
			Id:          user.Id,
			UserName:    user.UserName,
			UserAccount: user.UserAccount,
			AvatarUrl:   user.AvatarUrl,
			Phone:       user.Phone,
			Email:       user.Email,
			UserStatus:  user.UserStatus,
			Gender:      user.Gender,
			UserRole:    user.Role,
			Tags:        user.Tags,
			Profile:     user.Profile,
			CreateTime:  user.CreateTime.Format("2006-01-02"),
		},
	}, nil
}

func (s *PartnerService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*emptypb.Empty, error) {
	updateUser := &biz.UpdateUser{
		Id:        req.Data.Id,
		UserName:  req.Data.UserName,
		AvatarUrl: req.Data.AvatarUrl,
		Gender:    req.Data.Gender,
		Phone:     req.Data.Phone,
		Email:     req.Data.Email,
	}
	err := s.vc.ParamsValidate(updateUser)
	if err != nil {
		return nil, err
	}

	err = s.uc.UpdateUser(ctx, updateUser)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *PartnerService) UsersRecommend(ctx context.Context, req *v1.UsersRecommendReq) (*v1.UsersRecommendReply, error) {
	usersList, err := s.uc.UsersRecommend(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.UsersRecommendReply{
		Data: make([]*v1.User, 0, len(usersList)),
	}
	for _, item := range usersList {
		reply.Data = append(reply.Data, &v1.User{
			Id:          item.Id,
			UserName:    item.UserName,
			UserAccount: item.UserAccount,
			AvatarUrl:   item.AvatarUrl,
			Phone:       item.Phone,
			Email:       item.Email,
			UserStatus:  item.UserStatus,
			Gender:      item.Gender,
			UserRole:    item.Role,
			CreateTime:  item.CreateTime.String(),
			Tags:        item.Tags,
			Profile:     item.Profile,
		})
	}
	return reply, nil
}
