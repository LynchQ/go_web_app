package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	// dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"))
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		// fmt.Printf("connect DB failed, err:%v\n", err)
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) // 设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) // 设置连接池中的最大闲置连接数
	return
}

func Close() {
	_ = db.Close()
}
