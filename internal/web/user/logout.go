package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/oliver258/eagle/internal/web"
	"github.com/oliver258/eagle/pkg/log"
)

// Logout user logout
func Logout(c *gin.Context) {
	// 删除cookie信息
	session := web.GetCookieSession(c)
	session.Options = &sessions.Options{
		Domain: "",
		Path:   "/",
		MaxAge: -1,
	}
	err := session.Save(web.Request(c), web.ResponseWriter(c))
	if err != nil {
		log.Warnf("[user] logout save session err: %v", err)
		c.Abort()
		return
	}

	// 重定向得到原页面
	c.Redirect(http.StatusSeeOther, c.Request.Referer())
}
