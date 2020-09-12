package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/dto"
	"github.com/hsimao/go_gateway_demo/middleware"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

// AdminLogin godoc
// @Summary 管理員登入
// @Description 管理員登入
// @Tags 管理員接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOupt} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	out := &dto.AdminLoginOupt{Token: params.UserName}
	middleware.ResponseSuccess(c, out)
}
