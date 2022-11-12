package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gormtest/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(116.205.179.34:3306)/yundao?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	fmt.Println("err", err)

	db.AutoMigrate(&model.User{})
	entry, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-11-11 00:00:01", time.Local)
	fmt.Println("dd", entry)
	db.Save(&model.User{
		Name:         "张三",
		Level:        model.Master,
		ServiceCount: 20, YearsService: 10,
		PersonalProfile: "从业大师",
		Birthday:        time.Now(),
		TimeEntry:       entry,
	})

	var userss model.User
	db.Where(" id=1").Find(&userss)
	str, _ := json.Marshal(userss)
	fmt.Println(string(str))
	fmt.Println(&userss)
	db.Delete(&model.User{},"id=1")
	db.Model(&model.User{}).Where("name=?","张三1").Updates(&model.User{Name: "李四",ServiceCount: 100})
}
