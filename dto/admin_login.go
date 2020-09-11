package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/public"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,is_valid_username"`
	Password string `json:"password" form:"password" comment:"密碼" example:"123456" validate:"required"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
