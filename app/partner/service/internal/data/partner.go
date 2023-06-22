package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/pkg/util"
	"github.com/pkg/errors"
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

func (r *partnerRepo) AddTeam(ctx context.Context, team *biz.CreateTeam) error {
	et, err := time.ParseInLocation("2006-01-02 15:04:05", team.ExpireTime, time.Local)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to covert string to datatime: team(%v)", team))
	}
	t := &Team{}
	util.StructAssign(t, team)
	t.ExpireTime = et
	err = r.data.db.WithContext(ctx).Create(t).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to add a team: team(%v)", team))
	}
	return nil
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

func (r *partnerRepo) UpdateTeam(ctx context.Context, team *biz.UpdateTeam) error {
	t := &Team{}
	util.StructAssign(t, team)
	err := r.data.db.WithContext(ctx).Where("id = ? isDelete = 0", t.Id).Updates(t).Error
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
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("get team failed: teamId(%v)", teamId))
	}

	result := &biz.Team{}
	util.StructAssign(result, team)
	return result, nil
}

func (r *partnerRepo) GetTeamList(ctx context.Context, query *biz.TeamQuery, page, pageSize int32) ([]*biz.Team, error) {
	if page < 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 10
	}
	list := make([]*Team, 0)
	queryDB := r.data.db.WithContext(ctx).Where("isDelete = 0")
	// 链式查询
	for key, value := range util.PrintNonZeroFieldsAndValues(query) {
		queryDB = queryDB.Where(fmt.Sprintf("%s like ?", key), "%"+fmt.Sprintf("%v", value)+"%")
	}
	err := queryDB.Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to search team by query params: query(%v)", query))
	}

	var search []*biz.Team
	for _, item := range list {
		team := &biz.Team{}
		util.StructAssign(team, item)
		search = append(search, team)
	}
	return search, nil
}
