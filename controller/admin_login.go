package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/dao"
	"github.com/hsimao/go_gateway_demo/dto"
	"github.com/hsimao/go_gateway_demo/golang_common/lib"
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
		middleware.ResponseError(c, 2000, err)
		return
	}
	// 1. 從 params.UserName 取得管理員資料 admininfo
	// 2. admininfo.salt + params.Password sha256 => saltPassword
	// 3. saltPassword == admininfo.password
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	out := &dto.AdminLoginOupt{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}
