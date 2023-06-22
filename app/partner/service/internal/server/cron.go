package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redsync/redsync/v4"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/robfig/cron"
	_ "github.com/robfig/cron"
	"reflect"
	"time"
)

type CronJob struct {
	cron *cron.Cron
}

type Job struct {
	uc      *biz.UserUseCase
	redSync *redsync.Mutex
}

func NewCronJob(redSync *redsync.Mutex, uc *biz.UserUseCase) *CronJob {
	c := cron.New()
	j := &Job{uc: uc, redSync: redSync}
	st := reflect.TypeOf(j)
	sv := reflect.ValueOf(j)
	for i := 0; i < st.NumMethod(); i++ {
		method := st.Method(i)
		methodName := method.Name
		result := sv.MethodByName(methodName)
		// 每小时执行一次定时任务
		err := c.AddFunc("0 0 * * * *", result.Interface().(func()))
		if err != nil {
			log.Fatalf("cronjob init failed: %s", err.Error())
		}
		fmt.Println("Result:", result)
	}
	return &CronJob{
		cron: c,
	}
}

func (j *Job) UsersRecommend() {
	// 加锁
	if err := j.redSync.Lock(); err != nil {
		log.Errorf("get lock failed: %s", err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := j.uc.UsersRecommend(ctx, 1, 10)
	if err != nil {
		log.Errorf("cronjob (%s) error: %s", "UsersRecommend", err.Error())
	}
	// 释放锁
	if ok, err := j.redSync.Unlock(); !ok || err != nil {
		log.Errorf("release lock failed: %s", err.Error())
	}
}

func (s *CronJob) Start(_ context.Context) error {
	s.cron.Start()
	return nil
}

func (s *CronJob) Stop(_ context.Context) error {
	s.cron.Stop()
	return nil
}
