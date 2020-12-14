package test

import (
	"go-admin/models"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestDB(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@(192.168.168.150:3306)/go_admin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error(err)
	}
	db.LogMode(true)
	// u := models.User{Username: "test", Password: "test", CreateBy: 1, UpdateBy: 1,
	// 	Model: models.Model{CreateTime: time.Now(), UpdateTime: time.Now()}}
	// db.Create(&u)
	u := models.User{Model: models.Model{ID: 2}}
	db.Preload("Role").Find(&u)
	t.Logf("user:%+v", u)
	defer db.Close()
}
