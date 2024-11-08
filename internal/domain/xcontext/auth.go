package xcontext

import (
	"inbody-ocr-backend/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

var (
	authMemberUserKey = "AUTH:MEMBER"
	authAdminUserKey  = "AUTH:ADMIN"
	authOwnerUserKey  = "AUTH:OWNER"
)

func WithMemberUser(c *gin.Context, user entity.User) {
	c.Set(authMemberUserKey, user)
}

func MemberUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authMemberUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}

func WithAdminUser(c *gin.Context, user entity.User) {
	c.Set(authAdminUserKey, user)
}

func AdminUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authAdminUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}

func WithOwnerUser(c *gin.Context, user entity.User) {
	c.Set(authOwnerUserKey, user)
}

func OwnerUser(c *gin.Context) *entity.User {
	if v, ok := c.Get(authOwnerUserKey); ok {
		return v.(*entity.User)
	}
	return nil
}
