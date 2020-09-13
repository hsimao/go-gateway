package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/dto"
	"github.com/hsimao/go_gateway_demo/middleware"
	"github.com/hsimao/go_gateway_demo/public"
)

type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/admin_info", adminLogin.AdminInfo)
}

// AdminInfo godoc
// @Summary 管理員資料
// @Description 管理員資料
// @Tags 管理員接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminlogin *AdminController) AdminInfo(c *gin.Context) {

	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminInfoOutput{}
	err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 1. 讀取 sessionKey 對應 json 轉換為結構體
	// 2. 取出數據然後封裝輸出結構體
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.Name,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://cdn.iconscout.com/icon/free/png-512/avatar-370-456322.png",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}
