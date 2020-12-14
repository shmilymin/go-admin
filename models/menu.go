package models

import (
	u "go-admin/pkg/util"
	"time"
)

type Menus []Menu

// Menu
type Menu struct {
	Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	CreateBy uint64 `json:"createBy"`
	UpdateBy uint64 `json:"updateBy"`
}

// 指定表名
func (Menu) TableName() string {
	return "menu"
}

// 新增
func (m *Menu) Add() error {
	now := time.Now()
	m.CreateTime = now
	m.UpdateTime = now
	return DB.Create(&m).Error
}

// 根据Id查询
func (m *Menu) Get() error {
	return DB.First(&m).Error
}

// 修改
func (m *Menu) Update() error {
	m.UpdateTime = time.Now()
	return DB.Model(&m).Update(m).Error
}

// 获取列表
func (ms *Menus) List(e Menu, p u.Page) error {
	if p.Page == 0 || p.Limit == 0 {
		return DB.Find(&ms, e).Error
	} else {
		return DB.Limit(p.Limit).Offset((p.Page-1)*p.Limit).Find(&ms, e).Error
	}
}

// 获取总数
func (m *Menu) Count() (int, error) {
	var count int
	return count, DB.Model(Menu{}).Where(&m).Count(&count).Error
}

// 根据Id删除
func (m *Menu) Delete() error {
	return DB.Delete(&m).Error
}
