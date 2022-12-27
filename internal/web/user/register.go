package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oliver258/eagle/internal/model"
	"github.com/oliver258/eagle/internal/service"
	"github.com/oliver258/eagle/internal/web"
	"github.com/oliver258/eagle/pkg/errcode"
	"github.com/oliver258/eagle/pkg/flash"
	"github.com/oliver258/eagle/pkg/log"
)

// GetRegister register as a new user
func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register", gin.H{
		"title": "注册",
		"ctx":   c,
	})
}

// DoRegister submit register
func DoRegister(c *gin.Context) {
	log.Info("User Register function called.")
	var r RegisterRequest
	if err := c.Bind(&r); err != nil {
		web.Response(c, errcode.ErrInvalidParam, nil)
		return
	}

	u := model.UserBaseModel{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		web.Response(c, errcode.ErrValidation, nil)
		return
	}

	// Insert the user to the database.
	err := service.Svc.Users().Register(c, u.Username, u.Email, r.Password)
	if err != nil {
		web.Response(c, errcode.ErrInternalServer, nil)
		return
	}

	flash.SetMessage(c.Writer, "已发送激活链接,请检查您的邮箱。")

	// Show the user information.
	web.Response(c, nil, RegisterResponse{
		ID: u.ID,
	})
}
