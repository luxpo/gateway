package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/luxpo/gateway/public"
)

type AdminLoginInput struct {
	Username string `form:"username" json:"username" comment:"用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

type AdminLoginOutput struct {
	Token string `form:"token" json:"token" comment:"token"  validate:"" example:"token"`
}

func (params *AdminLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
