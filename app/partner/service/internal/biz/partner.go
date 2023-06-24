package biz

import (
	"context"
	"fmt"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"time"
)

type PartnerRepo interface {
	AddTeam(ctx context.Context, team *CreateTeam) (int32, error)
	AddUserTeam(ctx context.Context, userTeam *UserTeam) (int32, error)
	DeleteTeam(ctx context.Context, teamId int32) error
	DeleteUserTeam(ctx context.Context, teamId, userId int32) error
	UpdateTeam(ctx context.Context, team *UpdateTeam) error
	GetTeam(ctx context.Context, teamId int32) (*Team, error)
	GetTeamCountByUserId(ctx context.Context, userId int32) (int64, error)
	GetTeamList(ctx context.Context, query *TeamQuery, page, pageSizer int32) ([]*TeamList, error)
	GetUserTeamListByUserId(ctx context.Context, userId int32) ([]*UserTeam, error)
	GetUserTeamCountByTeamId(ctx context.Context, teamId int32) (int64, error)
	GetTeamNextManager(ctx context.Context, teamId, userId int32) (int32, error)
}

type PartnerRepoUseCase struct {
	repo PartnerRepo
	log  *log.Helper
	re   Recovery
	tm   Transaction
}

type Team struct {
	Id          int32  `validate:"omitempty,gte=0" comment:"id"`
	Name        string `validate:"omitempty,min=1,max=20" comment:"队伍名"`
	Description string `validate:"omitempty,max=512" comment:"描述"`
	MaxNum      int32  `validate:"omitempty,gt=1,lte=20" comment:"最大人数"`
	ExpireTime  string `validate:"omitempty,datetime" comment:"过期时间"`
	UserId      int32  `validate:"omitempty,gte=0" comment:"创建者用户id"`
	Status      int32  `validate:"omitempty,oneof=0 1 2" comment:"状态"`
	Password    string `validate:"omitempty,max=32" comment:"密码"`
}

type UserTeam struct {
	Id     int32
	UserId int32
	TeamId int32
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
//5. 如果 status 是加密状态，一定要有密码，且密码 <= 32 （有点复杂，AddTeam里实现）
//6. 超时时间 > 当前时间 （有点复杂，AddTeam里实现）
type CreateTeam struct {
	Team
	Name       string `validate:"required,min=1,max=20" comment:"队伍名"`
	MaxNum     int32  `validate:"required,gt=1,lte=20" comment:"最大人数"`
	Status     int32  `validate:"required,oneof=0 1 2" comment:"状态"`
	ExpireTime string `validate:"required,datetime" comment:"过期时间"`
	UserId     int32  `validate:"required,gte=0" comment:"创建者用户id"`
}

// AddTeam
//1. 校验用户最多创建 5 个队伍
//2. add
func (r *PartnerRepoUseCase) AddTeam(ctx context.Context, user *User, team *CreateTeam) error {
	err := r.addTeamParamsValidate(team)
	if err != nil {
		return err
	}

	count, err := r.repo.GetTeamCountByUserId(ctx, user.Id)
	if err != nil {
		return v1.ErrorUnknownError("未知错误: ", err.Error())
	}

	if count > 5 {
		return v1.ErrorAddTeamFailed("用户最多创建5个队伍")
	}

	err = r.tm.ExecTx(ctx, func(ctx context.Context) error {
		teamId, err := r.repo.AddTeam(ctx, team)
		if err != nil {
			return v1.ErrorAddTeamFailed("队伍创建失败: %s", err.Error())
		}

		_, err = r.repo.AddUserTeam(ctx, &UserTeam{
			UserId: user.Id,
			TeamId: teamId,
		})
		if err != nil {
			return v1.ErrorAddUserTeamFailed("队伍创建失败: %s", err.Error())
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *PartnerRepoUseCase) addTeamParamsValidate(team *CreateTeam) error {
	if team.Status == 2 && team.Password == "" {
		return v1.ErrorAddTeamFailed("密码不能为空")
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", team.ExpireTime)
	if err != nil {
		return v1.ErrorUnknownError(err.Error())
	}

	if parsedTime.Before(time.Now()) {
		return v1.ErrorAddTeamFailed("过期时间早于今天")
	}

	return nil
}

type DeleteTeam struct {
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) DeleteTeam(ctx context.Context, team *DeleteTeam) error {
	err := r.repo.DeleteTeam(ctx, team.Id)
	if err != nil {
		return v1.ErrorDeleteTeamFailed("队伍删除错误: %s", err.Error())
	}
	return nil
}

type UpdateTeam struct {
	Team
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) UpdateTeam(ctx context.Context, team *UpdateTeam, userId int32, isAdmin bool) error {
	teamInfo, err := r.repo.GetTeam(ctx, team.Id)
	if kerrors.IsNotFound(err) {
		return v1.ErrorUpdateTeamFailed("该队伍不存在")
	}
	if err != nil {
		return v1.ErrorUpdateTeamFailed("未知错误: %s", err.Error())
	}

	// 只有管理员或者队伍的创建者可以修改
	if !isAdmin && teamInfo.UserId != userId {
		return v1.ErrorPermissionDeny("只有管理员或者队伍的创建者可以修改")
	}
	//如果队伍状态改为加密，必须要有密码
	if team.Status == 2 && team.Password == "" {
		return v1.ErrorUpdateTeamFailed("密码不能为空")
	}
	err = r.repo.UpdateTeam(ctx, team)
	if err != nil {
		return v1.ErrorUpdateTeamFailed("队伍更新错误: %s", err.Error())
	}
	return nil
}

type GetTeam struct {
	Id int32 `validate:"required,gte=0" comment:"id"`
}

func (r *PartnerRepoUseCase) GetTeam(ctx context.Context, team *GetTeam) (*Team, error) {
	t, err := r.repo.GetTeam(ctx, team.Id)
	if err != nil {
		return nil, v1.ErrorGetTeamFailed("队伍获取错误: %s", err.Error())
	}
	return t, nil
}

type TeamQuery struct {
	Team
	searchText string
}

type TeamList struct {
	Team
	UserInfo *User
}

func (r *PartnerRepoUseCase) GetTeamList(ctx context.Context, team *TeamQuery, page, pageSize int32) ([]*TeamList, error) {
	list, err := r.repo.GetTeamList(ctx, team, page, pageSize)
	if err != nil {
		return nil, v1.ErrorGetTeamListFailed("队伍列表获取错误: %s", err.Error())
	}
	return list, nil
}

type JoinTeam struct {
	Team
	Id int32 `validate:"required,gte=0" comment:"id"`
}

// JoinTeam
//1. 用户最多加入 5 个队伍
//2. 队伍必须存在，只能加入未满、未过期的队伍
//3. 不能加入自己的队伍，不能重复加入已加入的队伍（幂等性）
//4. 禁止加入私有的队伍
//5. 如果加入的队伍是加密的，必须密码匹配才可以
//6. 新增队伍 - 用户关联信息
func (r *PartnerRepoUseCase) JoinTeam(ctx context.Context, team *JoinTeam, user *User) error {
	userTeamList, err := r.repo.GetUserTeamListByUserId(ctx, user.Id)
	if err != nil {
		return v1.ErrorJoinTeamFailed("")
	}

	//1. 用户最多加入 5 个队伍
	if len(userTeamList) > 5 {
		return v1.ErrorJoinTeamFailed("用户最多加入5个队伍")
	}

	//2. 队伍必须存在，只能加入未满、未过期的队伍
	teamInfo, err := r.repo.GetTeam(ctx, team.Id)
	if kerrors.IsNotFound(err) {
		return v1.ErrorJoinTeamFailed("该队伍不存在")
	}
	if err != nil {
		return v1.ErrorJoinTeamFailed("加入队伍出错: %s", err.Error())
	}

	// 只能加入未满的队伍
	userCount, err := r.repo.GetUserTeamCountByTeamId(ctx, team.Id)
	if err != nil {
		return v1.ErrorJoinTeamFailed("加入队伍出错: %s", err.Error())
	}
	if userCount > int64(teamInfo.MaxNum) {
		return v1.ErrorJoinTeamFailed("队伍已满")
	}

	// 只能加入未过期的队伍
	parsedTime, err := time.Parse("2006-01-02 15:04:05", teamInfo.ExpireTime)
	if err != nil {
		return v1.ErrorUnknownError(fmt.Sprintf("未知错误: %s", err.Error()))
	}

	if parsedTime.Before(time.Now()) {
		return v1.ErrorJoinTeamFailed("你不能加入一个已经过期的队伍")
	}

	//3. 不能加入自己的队伍，不能重复加入已加入的队伍（幂等性）
	for _, item := range userTeamList {
		if item.TeamId == team.Id {
			return v1.ErrorJoinTeamFailed("不能加入自己的队伍，不能重复加入已加入的队伍")
		}
	}

	//4. 禁止加入私有的队伍
	if teamInfo.Status == 1 {
		return v1.ErrorPermissionDeny("禁止加入私有的队伍")
	}

	//5. 如果加入的队伍是加密的，必须密码匹配才可以
	if teamInfo.Status == 2 && team.Password != teamInfo.Password {
		return v1.ErrorPermissionDeny("密码不匹配")
	}

	//6. 新增队伍 - 用户关联信息
	_, err = r.repo.AddUserTeam(ctx, &UserTeam{
		TeamId: team.Id,
		UserId: user.Id,
	})
	if teamInfo.Status == 2 && team.Password != teamInfo.Password {
		return v1.ErrorJoinTeamFailed("加入队伍出错: %s", err.Error())
	}

	return nil
}

type QuitTeam struct {
	Team
	Id int32 `validate:"required,gte=0" comment:"id"`
}

// QuitTeam
//1. 只剩一人，队伍解散
//
//2. 还有其他人
//
//1. 如果是队长退出队伍，权限转移给第二早加入的用户 —— 先来后到
//
//> 只用取 id 最小的 2 条数据
//
//2. 非队长，自己退出队伍

func (r *PartnerRepoUseCase) QuitTeam(ctx context.Context, team *QuitTeam, user *User) error {
	teamInfo, err := r.repo.GetTeam(ctx, team.Id)
	if kerrors.IsNotFound(err) {
		return v1.ErrorQuitTeamFailed("该队伍不存在")
	}
	if err != nil {
		return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
	}

	userCount, err := r.repo.GetUserTeamCountByTeamId(ctx, team.Id)
	if err != nil {
		return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
	}

	if userCount == 1 {
		//1. 只剩一人，队伍解散
		err = r.tm.ExecTx(ctx, func(ctx context.Context) error {
			err = r.repo.DeleteTeam(ctx, team.Id)
			if kerrors.IsNotFound(err) {
				return v1.ErrorQuitTeamFailed("该队伍不存在")
			}
			if err != nil {
				return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
			}

			err = r.repo.DeleteUserTeam(ctx, team.Id, user.Id)
			if kerrors.IsNotFound(err) {
				return v1.ErrorQuitTeamFailed("未曾加入过该队伍")
			}
			if err != nil {
				return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}

	//2. 还有其他人
	err = r.tm.ExecTx(ctx, func(ctx context.Context) error {
		//如果是队长退出队伍，权限转移给第二早加入的用户 —— 先来后到
		//只用取 id 最小的 2 条数据
		if teamInfo.UserId == user.Id {
			nextUser, err := r.repo.GetTeamNextManager(ctx, team.Id, user.Id)
			if err != nil {
				return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
			}

			updateTeam := &UpdateTeam{}
			updateTeam.Id = team.Id
			updateTeam.UserId = nextUser
			err = r.repo.UpdateTeam(ctx, updateTeam)
			if err != nil {
				return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
			}
		}

		//非队长，自己退出队伍
		err = r.repo.DeleteUserTeam(ctx, team.Id, user.Id)
		if kerrors.IsNotFound(err) {
			return v1.ErrorQuitTeamFailed("未曾加入过该队伍")
		}
		if err != nil {
			return v1.ErrorQuitTeamFailed("退出队伍出错: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
