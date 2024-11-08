package middleware

import (
	"fmt"
	"inbody-ocr-backend/internal/controller/render"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	tokenService service.TokenService
	userRepo     repository.UserRepository
}

func NewAPI(tokenService service.TokenService, userRepo repository.UserRepository) *API {
	return &API{
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

func (a *API) withUser(c *gin.Context) error {
	isValid, err := a.tokenService.TokenValid(c)
	if err != nil || !isValid {
		logging.Errorf(c, "withUser TokenValid %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	userID, orgID, err := a.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "withMember ExtractIDsFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		logging.Errorf(c, "withMember FindByID %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	if user.OrganizationID != orgID {
		logging.Errorf(c, "withMember OrgID mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("org_id mismatch")
	}

	xcontext.WithUser(c, user)

	return nil
}

func (a *API) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.withUser(c)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func (a *API) withMember(c *gin.Context) error {
	isValid, err := a.tokenService.TokenValid(c)
	if err != nil || !isValid {
		logging.Errorf(c, "withMember TokenValid %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	userID, orgID, err := a.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "withMember ExtractIDsFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		logging.Errorf(c, "withMember FindByID %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	if user.Role != "member" {
		logging.Errorf(c, "withMember Role mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("role mismatch")
	}

	if user.OrganizationID != orgID {
		logging.Errorf(c, "withMember OrgID mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("org_id mismatch")
	}

	xcontext.WithMemberUser(c, user)

	return nil
}

func (a *API) GuaranteeMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.withMember(c)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func (a *API) withAgminOROwner(c *gin.Context) error {
	isValid, err := a.tokenService.TokenValid(c)
	if err != nil || !isValid {
		logging.Errorf(c, "withAdminOROwner TokenValid %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	userID, orgID, err := a.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "withAdminOROwner ExtractIDsFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		logging.Errorf(c, "withAdminOROwner FindByID %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	if user.Role == "member" {
		logging.Errorf(c, "withAdminOROwner Role mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("role mismatch")
	}

	if user.OrganizationID != orgID {
		logging.Errorf(c, "withAdminOROwner OrgID mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("org_id mismatch")
	}

	// admin, ownerのどちらもアクセス可能なエンドポイントの場合、adminとして扱う
	xcontext.WithAdminUser(c, user)

	return nil
}

func (a *API) GuaranteeAdminOROwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.withAgminOROwner(c)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func (a *API) withAdmin(c *gin.Context) error {
	isValid, err := a.tokenService.TokenValid(c)
	if err != nil || !isValid {
		logging.Errorf(c, "withAdmin TokenValid %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	userID, orgID, err := a.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "withAdmin ExtractIDsFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		logging.Errorf(c, "withAdmin FindByID %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	if user.Role != "admin" {
		logging.Errorf(c, "withAdmin Role mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("role mismatch")
	}

	if user.OrganizationID != orgID {
		logging.Errorf(c, "withAdmin OrgID mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("org_id mismatch")
	}

	xcontext.WithAdminUser(c, user)

	return nil
}

func (a *API) GuaranteeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.withAdmin(c)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}

func (a *API) withOwner(c *gin.Context) error {
	isValid, err := a.tokenService.TokenValid(c)
	if err != nil || !isValid {
		logging.Errorf(c, "withOwner TokenValid %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	userID, orgID, err := a.tokenService.ExtractIDsFromContext(c)
	if err != nil {
		logging.Errorf(c, "withOwner ExtractIDsFromContext %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}

	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		logging.Errorf(c, "withOwner FindByID %v", err)
		render.ErrorJSON(c, err.Error(), http.StatusUnauthorized)
		return err
	}
	if user.Role != "owner" {
		logging.Errorf(c, "withOwner Role mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("role mismatch")
	}

	if user.OrganizationID != orgID {
		logging.Errorf(c, "withOwner OrgID mismatch")
		render.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized)
		return fmt.Errorf("org_id mismatch")
	}

	xcontext.WithOwnerUser(c, user)

	return nil
}

func (a *API) GuaranteeOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := a.withOwner(c)
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}
