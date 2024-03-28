package main

import (
    "fmt"
    "irisweb/config"
    "irisweb/model"
    "gorm.io/gorm"
    "irisweb"
    "gorm.io/driver/mysql"
)
var DB *gorm.DB
func main() {
    config.Server.Mysql.Database = "go_mysql"
    config.Server.Mysql.User = "root"
    config.Server.Mysql.Password = "123456"
    config.Server.Mysql.Host = "127.0.0.1"
    config.Server.Mysql.Port = 3306
    InitDB(&config.Server.Mysql)
    b := irisweb.New(config.Server.Server.Port, config.Server.Server.LogLevel)
    b.Serve()
}
func InitDB(setting *config.MysqlConfig) (*gorm.DB, error) {
    var db *gorm.DB
    var err error
    url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        setting.User, setting.Password, setting.Host, setting.Port, setting.Database)
        fmt.Printf(url)
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
