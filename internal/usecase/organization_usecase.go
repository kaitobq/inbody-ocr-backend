package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/infra/logger"
	"inbody-ocr-backend/internal/usecase/response"
)

type organizationUsecase struct {
	repo repository.OrganizationRepository
	membershipRepo repository.UserOrganizationMembershipRepository
	ulidService service.ULIDService
}

func NewOrganizationUsecase(repo repository.OrganizationRepository, membershipRepo repository.UserOrganizationMembershipRepository, ulidService service.ULIDService) OrganizationUsecase {
	return &organizationUsecase{
		repo: repo,
		membershipRepo: membershipRepo,
		ulidService: ulidService,
	}
}

func (uc *organizationUsecase) CreateOrganization(name string, founderID string) (*response.CreateOrganizationResponse, error) {
	id := uc.ulidService.GenerateULID()
	org := entity.Organization{
		ID: id,
		Name: name,
	}

	// ユーザーが作成したい組織名と同じ名前の組織に所属していないかチェック
	memberships, err := uc.membershipRepo.FindByUserID(founderID)
	if err != nil {
		logger.Error("CreateOrganization", "func", "FindByUserID()", "error", err.Error())
		return nil, err
	}
	if len(*memberships) != 0 {
		for _, membership := range *memberships {
			org, err := uc.repo.FindByID(membership.OrganizationID)
			if err != nil {
				logger.Error("CreateOrganization", "func", "FindByID()", "error", err.Error())
				return nil, err
			}

			if org.Name == name {
				logger.Info("User tryed to create organization with same name", "name", name)
				return nil, fmt.Errorf("cannot belong to multiple organizations with the same name [%s]", name)
			}
		}
	}

	organization, err := uc.repo.CreateOrganization(org)
	if err != nil {
		logger.Error("CreateOrganization", "func", "CreateOrganization()", "error", err.Error())
		return nil, err
	}

	id = uc.ulidService.GenerateULID()
	membership := entity.UserOrganizationMembership{
		ID: id,
		UserID: founderID,
		OrganizationID: organization.ID,
		Role: entity.OrganizationRoleOwner,
	}
	_, err = uc.membershipRepo.CreateMembership(membership)
	if err != nil {
		logger.Error("CreateOrganization", "func", "CreateMembership()", "error", err.Error())
		// 作成者が組織に所属できない場合はロールバック
		rollbackErr := uc.repo.DeleteOrganization(organization.ID)
		if rollbackErr != nil {
			logger.Error("CreateOrganization", "func", "DeleteOrganization()", "error", rollbackErr.Error())
			return nil, fmt.Errorf("failed to create membership and rollback organization creation: %w", err)
		}

		return nil, fmt.Errorf("failed to create membership: %w", err)
	}

	return response.NewCreateOrganizationResponse(organization.ID, organization.Name, organization.CreatedAt, organization.UpdatedAt)
}
