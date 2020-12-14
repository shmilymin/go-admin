package menu

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"go-admin/models"
	u "go-admin/pkg/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary 新增菜单
// @Tags 菜单
// @Produce json
// @Param menu body Menu true "menu"
// @Success 200
// @Failure 500
// @Router /api/menu [post]
func Add(c *gin.Context) {
	m := &models.Menu{}
	if err := c.Bind(&m); err != nil {
		u.Fail(c)
		return
	}
	// 校验菜单名是否存在
	ms := models.Menus{}
	if err := ms.List(models.Menu{Name: m.Name}, u.Page{}); err != nil {
		u.Fail(c)
		return
	}
	if len(ms) > 0 {
		u.FailMsg(c, "菜单名已存在")
		return
	}

	cid := uint64(jwt.ExtractClaims(c)["id"].(float64))
	m.CreateBy = cid
	m.UpdateBy = cid
	if err := m.Add(); err != nil {
		u.Fail(c)
		return
	} else {
		u.Ok(c)
		return
	}
}

// @Summary 根据id查询
// @Tags 菜单
// @Produce json
// @Param id path uint64 true "id"
// @Success 200
// @Failure 500
// @Router /api/menu/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		u.Fail(c)
		return
	}
	m := &models.Menu{Model: models.Model{ID: id}}
	if err := m.Get(); err != nil || gorm.IsRecordNotFoundError(err) {
		u.FailMsg(c, "菜单不存在")
		return
	} else {
		u.OkData(c, map[string]*models.Menu{"menu": m})
		return
	}
}

// @Summary 根据id修改
// @Tags 菜单
// @Produce  json
// @Param menu body Menu true "menu"
// @Success 200
// @Failure 500
// @Router /api/menu [put]
func Update(c *gin.Context) {
	m := &models.Menu{}
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
		u.Ok(c)
		return
	}
}

// @Summary 获取菜单列表
// @Tags 菜单
// @Produce json
// @Param page path int "页码"
// @Param limit path int "每页大小"
// @Success 200
// @Failure 500
// @Router /api/menu [get]
func List(c *gin.Context) {
	p := &u.Page{}
	c.Bind(&p)
	log.Printf("p:%+v", p)

	m := &models.Menu{}
	c.Bind(&m)
	log.Printf("menu:%+v", m)

	ms := &models.Menus{}
	ms.List(*m, *p)

	count, _ := m.Count()
	log.Printf("count:%d", count)
	u.OkData(c, map[string]interface{}{"list": ms, "count": count})
	return
}

// @Summary 根据id删除
// @Tags 菜单
// @Produce  json
// @Param id path uint64 true "id"
// @Success 200
// @Failure 500
// @Router /api/menu/{id} [delete]
func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		u.Fail(c)
		return
	}
	m := &models.Menu{Model: models.Model{ID: id}}
	if err := m.Delete(); err != nil {
		u.FailMsg(c, "删除失败")
		return
	} else {
		u.Ok(c)
		return
	}
}
