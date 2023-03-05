package data

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"time"

	"sd-point/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	sysLog "log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewSessionRepo, NewUserRepo)

// Data .
type Data struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger log.Logger
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	db, err := NewDatabase(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Debugf("failed to new database: %+v, error: %v", c.Database, err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

// NewRegistrar 服务注册发现的提供方法
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

func NewDatabase(driver, source string) (db *gorm.DB, err error) {

	newLogger := gormLogger.New(
		sysLog.New(os.Stdout, "\r\n", sysLog.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(source), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			log.Debugf("failed to init database, error: %v", err)
			return
		}
	default:
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Errorf("failed to connect database, error: %v", err)
		return nil, err
	}

	// 设置连接池
	// 空闲
	sqlDb.SetMaxIdleConns(50)
	// 打开
	sqlDb.SetMaxOpenConns(100)
	// 超时
	sqlDb.SetConnMaxLifetime(time.Second * 30)

	err = DBAutoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DBAutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
