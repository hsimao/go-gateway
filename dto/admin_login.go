package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/public"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,is_valid_username"` // 姓名
	Password string `json:"password" form:"password" comment:"密碼" example:"123456" validate:"required"`                  // 密碼
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOupt struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}
