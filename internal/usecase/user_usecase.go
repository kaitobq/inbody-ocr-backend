package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/domain/xerror"
	"inbody-ocr-backend/internal/usecase/response"
)

type userUsecase struct {
	repo                      repository.UserRepository
	orgRepo                   repository.OrganizationRepository
	measurementDateRepo       repository.MeasurementDateRepository
	userMeasurementStatusRepo repository.UserMeasurementStatusRepository
	tokenService              service.TokenService
	ulidService               service.ULIDService
}

func NewUserUsecase(repo repository.UserRepository, orgRepo repository.OrganizationRepository, measurementDateRepo repository.MeasurementDateRepository, userMeasurementStatusRepo repository.UserMeasurementStatusRepository, tokenService service.TokenService, ulidService service.ULIDService) UserUsecase {
	return &userUsecase{
		repo:                      repo,
		orgRepo:                   orgRepo,
		measurementDateRepo:       measurementDateRepo,
		userMeasurementStatusRepo: userMeasurementStatusRepo,
		tokenService:              tokenService,
		ulidService:               ulidService,
	}
}

// この関数はユーザが組織に参加するときに使用する
func (uc *userUsecase) CreateUser(name, email, password, orgID string) (*response.SignUpResponse, error) {
	exists, err := uc.repo.UserExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, xerror.ErrEmailAlreadyExists
	}

	hashedPassword, err := uc.repo.HashPassword(password)
	if err != nil {
		return nil, err
	}

	tx, err := uc.repo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			if rerr := tx.Rollback(); rerr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rerr)
			}
			panic(p)
		} else if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rerr)
			}
		}
	}()

	id := uc.ulidService.GenerateULID()
	user := &entity.User{
		ID:             id,
		Name:           name,
		Email:          email,
		Password:       hashedPassword,
		OrganizationID: orgID,
		Role:           entity.OrganizationRoleMember,
	}

	user, err = uc.repo.CreateUserWithTx(tx, *user)
	if err != nil {
		return nil, err
	}

	// 組織に測定日があればstatusを作成する
	measurementDates, err := uc.measurementDateRepo.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	if len(measurementDates) > 0 {
		ent := entity.UserMeasurementStatus{
			ID:                uc.ulidService.GenerateULID(),
			UserID:            user.ID,
			MeasurementDateID: measurementDates[0].ID, // 最新
			HasRegistered:     false,
		}
		err = uc.userMeasurementStatusRepo.CreateUserMeasurementStatusWithTx(tx, ent)
	}

	token, err := uc.tokenService.GenerateTokenFromID(user.ID, user.OrganizationID)
	if err != nil {
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return response.NewSignUpResponse(token, exp, *user)
}

func (uc *userUsecase) SignIn(email, password string) (*response.SignInResponse, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = uc.repo.ComparePassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	token, err := uc.tokenService.GenerateTokenFromID(user.ID, user.OrganizationID)
	if err != nil {
		return nil, err
	}

	exp, err := uc.tokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewSignInResponse(token, exp, *user)
}

func (uc *userUsecase) GetOwnInfo(user entity.User) (*response.GetOwnInfoResponse, error) {
	return response.NewGetOwnInfoResponse(user)
}
