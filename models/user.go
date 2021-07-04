package models

import (
	u "go-admin/pkg/util"
	"time"
)

type Users []User

// User
type User struct {
	Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Gender   int    `json:"gender" gorm:"default:3"`
	CreateBy uint64 `json:"createBy"`
	UpdateBy uint64 `json:"updateBy"`

	Role []Role `json:"role" gorm:"many2many:user_role"`
}

// 指定表名
func (User) TableName() string {
	return "user"
}

// 新增
func (m *User) Add() error {
	now := time.Now()
	m.CreateTime = now
	m.UpdateTime = now
	return DB.Create(&m).Error
}

// 根据Id查询
func (m *User) Get() error {
	return DB.Preload("Role").First(&m).Error
}

// 修改
func (m *User) Update() error {
	m.UpdateTime = time.Now()
	return DB.Model(&m).Update(m).Error
}

// 获取列表
func (ms *Users) List(e User, p u.Page) error {
	if p.Page == 0 || p.Limit == 0 {
		return DB.Preload("Role").Find(&ms, e).Error
	} else {
		return DB.Preload("Role").Limit(p.Limit).Offset((p.Page-1)*p.Limit).Find(&ms, e).Error
	}
}

// 获取总数
func (m *User) Count() (int, error) {
	var count int
	return count, DB.Model(User{}).Where(&m).Count(&count).Error
}

// 根据Id删除
func (m *User) Delete() error {
	return DB.Delete(&m).Error
}
