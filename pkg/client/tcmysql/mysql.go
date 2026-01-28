package tcmysql

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/prometheus"
)

func InitMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // 单数表名
			NameReplacer:  strings.NewReplacer("PID", "pid"), // 缩写字段如何映射列名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键
		NowFunc: func() time.Time {
			return time.Now().Local() // 统一 DB 中存储的时间
		},
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(20)                // 缓存空闲连接
	sqlDB.SetMaxOpenConns(100)               // 最大连接数
	sqlDB.SetConnMaxLifetime(24 * time.Hour) // 连接最大生命周期

	registerMetrics(db)

	return db
}

// 监控 DB 状态: QPS,活跃连接,慢查询趋势,线程数
func registerMetrics(db *gorm.DB) {
	_ = db.Use(prometheus.New(prometheus.Config{
		DBName:          "db_data-front",
		RefreshInterval: 10,
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				Prefix:        "gorm_status_",
				VariableNames: []string{"Threads_running"},
			},
		},
	}))
}
