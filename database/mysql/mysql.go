package mysql

import (
	"context"
	"fmt"
	"time"

	"go_backend/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL(ctx context.Context) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.dbName"))

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   viper.GetString("mysql.tablePrefix"),
		},
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().In(location)
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("mysql.maxLifetime")) * time.Hour)

	if err := db.WithContext(ctx).AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
