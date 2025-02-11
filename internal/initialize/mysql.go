package initialize

import (
	"fmt"
	"time"

	"github.com/anonydev/e-commerce-api/global"
	"github.com/anonydev/e-commerce-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	// Connect to MySQL database
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	fmt.Printf("MySQL DSN: %s\n", s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect database")
	global.Logger.Info("Successfully connected to MySQL database")
	global.Mdb = db

	// Set MySQL pool
	SetPool()

	// migrate tables
	migrateTables()
}

func checkErrorPanic(err error, errMessage string) {
	if err != nil {
		global.Logger.Error(errMessage, zap.Error(err))
		panic(err)
	}
}

func SetPool() {
	// Set MySQL pool
	m := global.Config.MySQL
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Failed to set MySQL pool: %s::", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	// Limits Number of connectors in the pool
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	// After connect to the database, the maximum time that the connection is reused
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	// Migrate tables
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.ROLE{},
	)
	if err != nil {
		fmt.Println("Failed to migrate tables")
	}
}
