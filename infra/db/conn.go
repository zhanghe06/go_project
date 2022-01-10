package db

import (
	"database/sql"
	"fmt"
	"go_project/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"sync"
	"time"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

func connect() (*gorm.DB, error) {
	var err error
	if db != nil {
		return db, err
	}
	conf := config.NewConfig()
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.DbHost,
		conf.Mysql.DbPort,
		conf.Mysql.DbName,
		conf.Mysql.Charset,
		true,
		"Local",
	)

	// 超时配置
	if conf.Mysql.Timeout != "" {
		dsn = strings.Join([]string{dsn, conf.Mysql.Timeout}, "&")
	}
	if conf.Mysql.TimeoutRead != "" {
		dsn = strings.Join([]string{dsn, conf.Mysql.TimeoutRead}, "&")
	}
	if conf.Mysql.TimeoutWrite != "" {
		dsn = strings.Join([]string{dsn, conf.Mysql.TimeoutWrite}, "&")
	}
	// 打印慢查询
	//slowLogger := logger.New(
	//	//将标准输出作为Writer
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//
	//	logger.Config{
	//		//设定慢查询时间阈值为1ms
	//		SlowThreshold: 1 * time.Microsecond,
	//		//设置日志级别，只有Warn和Info级别会输出慢查询日志
	//		LogLevel: logger.Warn,
	//	},
	//)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
		// 打印慢查询
		// Logger: slowLogger,
	})
	if err != nil {
		return db, err
	}

	// 连接池配置
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		return db, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 500)
	// maxBadConnRetries 默认重试2次

	// 检测连接
	err = sqlDB.Ping()

	return db, err
}

func NewDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		db, err = connect()
		if err != nil {
			panic(err)
		}
	})
	return db
}
