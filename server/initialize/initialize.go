package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
	"ve-blog-go/server/global"
)

func Gorm() {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GVA_CONFIG.MysqlConfig.Username,
		global.GVA_CONFIG.MysqlConfig.Password,
		global.GVA_CONFIG.MysqlConfig.Host,
		global.GVA_CONFIG.MysqlConfig.Port,
		global.GVA_CONFIG.MysqlConfig.Dbname)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Info),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("GORM 数据库连接失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("SQL 数据库连接失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.GVA_LOG.Info("Mysql 数据库连接成功")
	global.GVA_DB = db
}

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("Redis 连接成功: " + pong)
		global.GVA_REDIS = client
	}
}
