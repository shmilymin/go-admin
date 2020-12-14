package mid

import (
	"errors"
	"go-admin/models"
	u "go-admin/pkg/util"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	identityKey = "username"
	Auth        *jwt.GinJWTMiddleware
	err         error
)

func init() {
	log.Println("init jwt")
	Auth, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     36 * time.Hour,
		MaxRefresh:  36 * time.Hour,
		IdentityKey: identityKey,
		// 登录时调用，可将载荷添加到token中
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			log.Printf("jwt PayloadFunc data:%+v", data)
			if user, ok := data.(models.User); ok {
				return jwt.MapClaims{
					"id":        user.ID,
					identityKey: user.Username,
				}
			}
			return jwt.MapClaims{}
		},
		// 验证登录状态
		IdentityHandler: func(c *gin.Context) interface{} {
			username := jwt.ExtractClaims(c)[identityKey]
			log.Printf("jwt IdentityHandler username:%s", username)
			return username
		},
		// 验证登录
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user := &models.User{}
			if err := c.ShouldBind(user); err != nil {
				return "", err
			}
			log.Printf("jwt login user:%+v", user)
			user.Password = u.EncodeMD5(user.Password)
			users := models.Users{}
			if err := users.List(*user, u.Page{}); err != nil || len(users) == 0 {
				return nil, errors.New("密码错误")
			}
			return users[0], nil
		},
		// 鉴权成功后执行
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		// 登录成功的回调函数
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			u.OkData(c, map[string]string{
				"token":  token,
				"expire": expire.Format("2006-01-02 15:04:05"),
			})
		},
		// 登录失效时的回调函数
		Unauthorized: func(c *gin.Context, code int, msg string) {
			log.Printf("jwt Unauthorized code:%d msg:%s", code, msg)
			u.CodeMsg(c, 886, msg)
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,

		// Optionally return the token as a cookie
		SendCookie: true,
	})
}
