package models

import (
	u "go-admin/pkg/util"
	"time"
)

type Roles []Role

// Role
type Role struct {
	Model
	Name     string `json:"name"`
	CreateBy uint64 `json:"createBy"`
	UpdateBy uint64 `json:"updateBy"`

	Menu []Menu `json:"menu" gorm:"many2many:role_menu"`
}

// 指定表名
func (Role) TableName() string {
	return "role"
}

// 新增
func (m *Role) Add() error {
	now := time.Now()
	m.CreateTime = now
	m.UpdateTime = now
	return DB.Create(&m).Error
}

// 根据Id查询
func (m *Role) Get() error {
	return DB.Preload("Menu").First(&m).Error
}

// 修改
func (m *Role) Update() error {
	m.UpdateTime = time.Now()
	return DB.Model(&m).Update(m).Error
}

// 获取列表
func (ms *Roles) List(e Role, p u.Page) error {
	if p.Page == 0 || p.Limit == 0 {
		return DB.Preload("Menu").Find(&ms, e).Error
	} else {
		return DB.Preload("Menu").Limit(p.Limit).Offset((p.Page-1)*p.Limit).Find(&ms, e).Error
	}
}

// 获取总数
func (m *Role) Count() (int, error) {
	var count int
	return count, DB.Model(Role{}).Where(&m).Count(&count).Error
}

// 根据Id删除
func (m *Role) Delete() error {
	return DB.Delete(&m).Error
}
