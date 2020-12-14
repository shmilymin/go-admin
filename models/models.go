package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID         uint64    `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@(192.168.168.150:3306)/go_admin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 设置连接池中的最大闲置连接数
	DB.DB().SetMaxIdleConns(10)
	// 设置数据库的最大连接数量
	DB.DB().SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	DB.DB().SetConnMaxLifetime(time.Hour)
	// 显示详细日志
	DB.LogMode(true)
}
