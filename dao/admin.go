package dao

import (
	"errors"
	"time"

	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/hsimao/go_gateway_demo/dto"
	"github.com/hsimao/go_gateway_demo/public"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主鍵"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理員用戶名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"Salt"`
	Password  string    `json:"password" gorm:"column:password" description:"密碼"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新時間"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"新增時間"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否刪除"`
}

// 資料庫 table 名稱
func (t *Admin) TableName() string {
	return "gateway_admin"
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := t.Find(c, tx, (&Admin{UserName: param.UserName, IsDelete: 0}))
	if err != nil {
		return nil, errors.New("用戶不存在")
	}

	saltPassword := public.GenSaltPassword(adminInfo.Salt, param.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密碼錯誤, 請重新輸入")
	}

	return adminInfo, nil
}

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
