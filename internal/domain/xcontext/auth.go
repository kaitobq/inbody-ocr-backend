package xcontext

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

var (
	authUserKey       = "AUTH:USER"   // ロール区分なし
	authMemberUserKey = "AUTH:MEMBER" // ロール区分：member
	authAdminUserKey  = "AUTH:ADMIN"  // ロール区分：admin
	authOwnerUserKey  = "AUTH:OWNER"  // ロール区分：owner
)

func WithUser(c *gin.Context, user *entity.User) {
	c.Set(authUserKey, user)
}

func User(c *gin.Context) *entity.User {
	if v, ok := c.Get(authUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}

func WithMemberUser(c *gin.Context, user *entity.User) {
	c.Set(authMemberUserKey, user)
}

func MemberUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authMemberUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}

func WithAdminUser(c *gin.Context, user *entity.User) {
	c.Set(authAdminUserKey, user)
}

func AdminUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authAdminUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}

func WithOwnerUser(c *gin.Context, user *entity.User) {
	c.Set(authOwnerUserKey, user)
}

func OwnerUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authOwnerUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}
