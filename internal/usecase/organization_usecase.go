package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
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
		return nil, err
	}
	if len(*memberships) != 0 {
		for _, membership := range *memberships {
			org, err := uc.repo.FindByID(membership.OrganizationID)
			if err != nil {
				return nil, err
			}

			if org.Name == name {
				return nil, fmt.Errorf("cannot belong to multiple organizations with the same name [%s]", name)
			}
		}
	}

	organization, err := uc.repo.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	role := "owner"
	err = uc.membershipRepo.CreateMembership(founderID, organization.ID, role)
	if err != nil {
		// 作成者が組織に所属できない場合はロールバック
		err = uc.repo.DeleteOrganization(organization.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to create membership and rollback organization creation: %w", err)
		}

		return nil, fmt.Errorf("failed to create membership: %w", err)
	}

	return response.NewCreateOrganizationResponse(organization.ID, organization.Name, organization.CreatedAt, organization.UpdatedAt)
}
