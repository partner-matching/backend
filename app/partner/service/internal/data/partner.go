package data

import (
	"context"
	"fmt"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

var _ biz.PartnerRepo = (*partnerRepo)(nil)

type partnerRepo struct {
	data *Data
	log  *log.Helper
}

func NewPartnerRepo(data *Data, logger log.Logger) biz.PartnerRepo {
	return &partnerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user/data/partner")),
	}
}

func (r *partnerRepo) AddTeam(ctx context.Context, team *biz.CreateTeam) (int32, error) {
	et, err := time.ParseInLocation("2006-01-02 15:04:05", team.ExpireTime, time.Local)
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to covert string to datatime: team(%v)", team))
	}
	t := &Team{}
	util.StructAssign(t, team)
	t.ExpireTime = et
	err = r.data.db.WithContext(ctx).Create(t).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to add a team: team(%v)", team))
	}
	return t.Id, nil
}

func (r *partnerRepo) AddUserTeam(ctx context.Context, userTeam *biz.UserTeam) (int32, error) {
	t := &UserTeam{}
	util.StructAssign(t, userTeam)
	t.JoinTime = time.Now()
	err := r.data.db.WithContext(ctx).Create(t).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to add user team: userTeam(%v)", userTeam))
	}
	return t.Id, nil
}

func (r *partnerRepo) DeleteTeam(ctx context.Context, teamId int32) error {
	team := &Team{
		Id:       teamId,
		IsDelete: 1,
	}
	err := r.data.db.WithContext(ctx).Where("id = ? and isDelete = 0", teamId).Delete(team).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to delete team: teamId(%v)", teamId))
	}
	return nil
}

func (r *partnerRepo) UpdateTeam(ctx context.Context, team *biz.UpdateTeam, userId int32, isAdmin bool) error {
	t := &Team{}
	util.StructAssign(t, team)
	queryDB := r.data.db.WithContext(ctx).Where("isDelete = 0")
	//只有管理员或者队伍的创建者可以修改
	switch isAdmin {
	case false:
		queryDB.Where("userId = ?", userId)
	}
	err := r.data.db.WithContext(ctx).Where("id = ?", t.Id).Updates(t).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to update team: team(%v)", team))
	}
	return nil
}

func (r *partnerRepo) GetTeam(ctx context.Context, teamId int32) (*biz.Team, error) {
	team := &Team{
		Id: teamId,
	}
	err := r.data.db.WithContext(ctx).Where("id = ? and isDelete = 0", teamId).First(team).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kerrors.NotFound("team is not found from db", fmt.Sprintln(teamId))
	}
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("get team failed: teamId(%v)", teamId))
	}

	result := &biz.Team{}
	util.StructAssign(result, team)
	return result, nil
}

func (r *partnerRepo) GetTeamCountByUserId(ctx context.Context, userId int32) (int64, error) {
	var count int64
	team := &Team{}
	err := r.data.db.WithContext(ctx).Model(team).Where("userId = ? and isDelete = 0", userId).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("get team count by user id failed: userId(%v)", userId))
	}
	return count, nil
}

// GetTeamList 分页展示队伍列表，根据名称、最大人数等搜索队伍  P0，信息流中不展示已过期的队伍
//
//1. 从请求参数中取出队伍名称等查询条件，如果存在则作为查询条件
//2. 不展示已过期的队伍（根据过期时间筛选）
//3. 可以通过某个**关键词**同时对名称和描述查询
//4. 关联查询已加入队伍的用户信息
//5. **关联查询已加入队伍的用户信息（可能会很耗费性能，建议大家用自己写 SQL 的方式实现）**
func (r *partnerRepo) GetTeamList(ctx context.Context, query *biz.TeamQuery, page, pageSize int32) ([]*biz.TeamList, error) {
	if page < 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	list := make([]*Team, 0)
	queryDB := r.data.db.WithContext(ctx)
	// 链式查询
	// 1. 从请求参数中取出队伍名称等查询条件，如果存在则作为查询条件
	for key, value := range util.PrintNonZeroFieldsAndValues(query) {
		// 首字母小写
		key = strings.ToLower(string(key[0])) + key[1:]
		value := fmt.Sprintf("%v", value)
		switch key {
		case "id", "maxNum", "userId", "status":
			queryDB = queryDB.Where(fmt.Sprintf("%s = ?", key), value)
		case "searchText":
			// 3. 可以通过某个**关键词**同时对名称和描述查询
			queryDB = queryDB.Where("name like ? or description like ? ", "%"+value+"%", "%"+value+"%")
		default:
			queryDB = queryDB.Where(fmt.Sprintf("%s like ?", key), "%"+fmt.Sprintf("%v", value)+"%")
		}
	}
	// 2. 不展示已过期的队伍（根据过期时间筛选）
	queryDB = queryDB.Where("expireTime is null or expireTime >= ? and isDelete = 0", time.Now())
	err := queryDB.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to search team by query params: query(%v)", query))
	}

	var search []*biz.TeamList
	for _, item := range list {
		team := &biz.TeamList{}
		util.StructAssign(team, item)

		// 获取创建者的信息
		userInfo := &User{}
		err = r.data.db.WithContext(ctx).Where("id = ? and isDelete = 0", item.UserId).First(userInfo).Error
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("get user failed: userId(%v)", item.UserId))
		}

		team.UserInfo = &biz.User{}
		util.StructAssign(team.UserInfo, userInfo)

		search = append(search, team)
	}

	return search, nil
}

func (r *partnerRepo) GetUserTeamListByUserId(ctx context.Context, userId int32) ([]*biz.UserTeam, error) {
	userTeamList := make([]*Team, 0)
	err := r.data.db.WithContext(ctx).Model(&UserTeam{}).Where("userId = ? and isDelete = 0", userId).Find(userTeamList).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("get user team by user id failed: userId(%v)", userId))
	}

	result := make([]*biz.UserTeam, 0)
	for _, item := range userTeamList {
		userTeam := &biz.UserTeam{}
		util.StructAssign(userTeam, item)
		result = append(result, userTeam)
	}
	return result, nil
}
