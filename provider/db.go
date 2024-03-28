package provider
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"irisweb/config"
	"irisweb/model"
)

var defaultDB *gorm.DB

func GetDefaultDB() *gorm.DB {
	if defaultDB == nil {
		config.Server.Mysql.Database = "go_mysql"
		config.Server.Mysql.User = "root"
		config.Server.Mysql.Password = "123456"
		config.Server.Mysql.Host = "127.0.0.1"
		config.Server.Mysql.Port = 3306
		InitDB(&config.Server.Mysql)
		db, err := InitDB(&config.Server.Mysql)
		if err != nil {
			fmt.Println("Failed To Connect Database: ", err.Error())
		}

		defaultDB = db
	}

	return defaultDB
}
func InitDB(setting *config.MysqlConfig) (*gorm.DB, error) {
    var db *gorm.DB
    var err error
    url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        setting.User, setting.Password, setting.Host, setting.Port, setting.Database)
    // setting.Url = url
    db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    sqlDB.SetMaxIdleConns(1000)
    sqlDB.SetMaxOpenConns(100000)
    sqlDB.SetConnMaxLifetime(-1)

    db.AutoMigrate(&model.User{}, &model.Category{}, &model.Goods{})


    return db, err
}