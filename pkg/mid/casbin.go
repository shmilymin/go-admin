package mid

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	u "go-admin/pkg/util"
	"log"
)

var E *casbin.Enforcer
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "admin"
`

// 初始化权限
func init() {
	log.Println("注入casbin")
	model, _ := model.NewModelFromString(text)
	E, _ = casbin.NewEnforcer(model)

	// 加载casbin策略数据，包括角色权限数据、用户角色数据
	LoadAllRolePolicy()
	LoadAllUserPolicy()
}

// 权限校验
func CasbinMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		if b, err := E.Enforce(jwt.ExtractClaims(c)[identityKey], c.Request.URL.Path, c.Request.Method); err != nil {
			u.CodeMsg(c, 401, "登录用户 校验权限失败")
			c.Abort()
			return
		} else if !b {
			u.CodeMsg(c, 401, "登录用户 没有权限")
			c.Abort()
			return
		}
		c.Next()
	}
}

// 加载全部角色权限
func LoadAllRolePolicy() error {
	ms := models.Roles{}
	if err := ms.List(models.Role{}, u.Page{}); err != nil {
		return err
	}
	for _, role := range ms {
		if err := LoadRolePolicy(role.ID); err != nil {
			return err
		}
	}
	return nil
}

// 根据角色id加载角色权限
func LoadRolePolicy(id uint64) error {
	role := models.Role{}
	role.ID = id
	if err := role.Get(); err != nil {
		return err
	}
	E.DeleteRole(role.Name)
	for _, menu := range role.Menu {
		if menu.Path == "" || menu.Method == "" {
			continue
		}
		E.AddPermissionForUser(role.Name, menu.Path, menu.Method)
	}
	return nil
}

// 加载全部用户角色策略
func LoadAllUserPolicy() error {
	ms := models.Users{}
	if err := ms.List(models.User{}, u.Page{}); err != nil {
		return err
	}
	for _, user := range ms {
		if err := LoadUserPolicy(user.ID); err != nil {
			return err
		}
	}
	log.Printf("用户角色关系:%+v", E.GetGroupingPolicy())
	return nil
}

// 根据用户id加载用户角色策略
func LoadUserPolicy(id uint64) error {
	user := models.User{}
	user.ID = id
	if err := user.Get(); err != nil {
		return err
	}
	E.DeleteRolesForUser(user.Username)
	for _, role := range user.Role {
		E.AddRoleForUser(user.Username, role.Name)
	}
	log.Printf("更新用户角色关系:%+v", E.GetGroupingPolicy())
	return nil
}
