package data

import (
	"os"
	"sd-point/internal/biz"
	"sd-point/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	sysLog "log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPointRepo)

// Data .
type Data struct {
	db     *gorm.DB
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

	return &Data{
		db: db,
	}, cleanup, nil
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
	return db.AutoMigrate(&biz.Point{})
}
