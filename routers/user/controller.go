package user

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-admin/models"
	"go-admin/pkg/mid"
	u "go-admin/pkg/util"
	"log"
	"strconv"
)

// @Summary 新增用户
// @Tags 用户
// @Produce json
// @Param user body User true "user"
// @Success 200
// @Failure 500
// @Router /api/user [post]
func Add(c *gin.Context) {
	m := &models.User{}
	if err := c.Bind(&m); err != nil {
		u.Fail(c)
		return
	}
	// 校验用户名是否存在
	ms := models.Users{}
	if err := ms.List(models.User{Username: m.Username}, u.Page{}); err != nil {
		u.Fail(c)
		return
	}
	if len(ms) > 0 {
		u.FailMsg(c, "用户名已存在")
		return
	}
	cid := uint64(jwt.ExtractClaims(c)["id"].(float64))
	m.CreateBy = cid
	m.UpdateBy = cid
	if err := m.Add(); err != nil {
		u.Fail(c)
		return
	} else {
		mid.LoadUserPolicy(m.ID)
		u.Ok(c)
		return
	}
}

// @Summary 根据id查询
// @Tags 用户
// @Produce json
// @Param id path uint64 true "id"
// @Success 200
// @Failure 500
// @Router /api/user/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		u.Fail(c)
		return
	}
	m := &models.User{Model: models.Model{ID: id}}
	if err := m.Get(); err != nil || gorm.IsRecordNotFoundError(err) {
		u.FailMsg(c, "用户不存在")
		return
	} else {
		u.OkData(c, map[string]*models.User{"user": m})
		return
	}
}

// @Summary 根据id修改
// @Tags 用户
// @Produce  json
// @Param user body User true "user"
// @Success 200
// @Failure 500
// @Router /api/user [put]
func Update(c *gin.Context) {
	m := &models.User{}
	if err := c.Bind(&m); err != nil {
		u.Fail(c)
		return
	}
	cid := uint64(jwt.ExtractClaims(c)["id"].(float64))
	m.UpdateBy = cid
	if err := m.Update(); err != nil {
		u.FailMsg(c, "修改失败")
		return
	} else {
		mid.LoadUserPolicy(m.ID)
		u.Ok(c)
		return
	}
}

// @Summary 获取用户列表
// @Tags 用户
// @Produce json
// @Param page path int "页码"
// @Param limit path int "每页大小"
// @Success 200
// @Failure 500
// @Router /api/user [get]
func List(c *gin.Context) {
	p := &u.Page{}
	c.Bind(&p)
	log.Printf("user list p:%+v", p)

	m := &models.User{}
	c.Bind(&m)
	log.Printf("user list user:%+v", m)

	ms := &models.Users{}
	ms.List(*m, *p)

	count, _ := m.Count()
	log.Printf("user list count:%d", count)
	u.OkData(c, map[string]interface{}{"list": ms, "count": count})
	return
}

// @Summary 根据id删除
// @Tags 用户
// @Produce  json
// @Param id path uint64 true "id"
// @Success 200
// @Failure 500
// @Router /api/user/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		u.Fail(c)
		return
	}
	m := &models.User{Model: models.Model{ID: id}}
	if err := m.Delete(); err != nil {
		u.FailMsg(c, "删除失败")
		return
	} else {
		mid.E.DeleteUser(m.Username)
		u.Ok(c)
		return
	}
}
