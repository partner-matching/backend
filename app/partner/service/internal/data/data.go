package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/partner-matching/backend/app/partner/service/internal/biz"
	"github.com/partner-matching/backend/app/partner/service/internal/conf"
	"gopkg.in/boj/redistore.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"runtime"
	"time"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewTransaction, NewRedis, NewRecovery, NewUserRepo, NewAuthRepo, NewSession)

type Data struct {
	log          *log.Helper
	db           *gorm.DB
	redisCli     redis.Cmdable
	conf         *conf.UserConstant
	sessionStore *redistore.RediStore
}

type contextTxKey struct{}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func (d *Data) GroupRecover(ctx context.Context, fn func(ctx context.Context) error) func() error {
	return func() error {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10)
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				d.log.Errorf("%v: %s\n", rerr, buf)
			}
		}()
		return fn(ctx)
	}
}

func (d *Data) Recover(ctx context.Context, fn func(ctx context.Context)) func() {
	return func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10)
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				d.log.Errorf("%v: %s\n", rerr, buf)
			}
		}()
		fn(ctx)
	}
}

func NewRecovery(d *Data) biz.Recovery {
	return d
}

// NewDB 初始化mysql
func NewDB(conf *conf.Data) *gorm.DB {
	l := log.NewHelper(log.With(log.GetLogger(), "module", "user/data/mysql"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}
	return db
}

// NewRedis redis
func NewRedis(conf *conf.Data) redis.Cmdable {
	l := log.NewHelper(log.With(log.GetLogger(), "module", "user/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr + conf.Redis.Port,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
		Password:     conf.Redis.Password,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		l.Fatalf("redis connect error: %v", err)
	}
	return client
}

// NewSession 初始化session
func NewSession(data *conf.Data, constant *conf.UserConstant) *redistore.RediStore {
	l := log.NewHelper(log.With(log.GetLogger(), "module", "user/data/session"))
	store, err := redistore.NewRediStore(10, "tcp", data.Redis.Addr+data.Redis.Port, "", []byte("secret-key"))
	if err != nil {
		l.Fatalf("new sessions error: %v", err)
	}
	store.SetMaxAge(int(constant.SessionTimeout))
	return store
}

func NewData(db *gorm.DB, redisCmd redis.Cmdable, logger log.Logger, conf *conf.UserConstant, sessionStore *redistore.RediStore) (*Data, func(), error) {
	l := log.NewHelper(log.With(log.GetLogger(), "module", "user/data/new-data"))

	d := &Data{
		log:          log.NewHelper(log.With(logger, "module", "creation/data")),
		db:           db,
		redisCli:     redisCmd,
		conf:         conf,
		sessionStore: sessionStore,
	}
	return d, func() {
		l.Info("closing the data resources")

		sqlDB, err := db.DB()
		if err != nil {
			l.Errorf("close db err: %v", err.Error())
		}

		err = sqlDB.Close()
		if err != nil {
			l.Errorf("close db err: %v", err.Error())
		}

		err = redisCmd.(*redis.Client).Close()
		if err != nil {
			l.Errorf("close redis err: %v", err.Error())
		}

		err = sessionStore.Close()
		if err != nil {
			l.Errorf("close session err: %v", err.Error())
		}
	}, nil
}
