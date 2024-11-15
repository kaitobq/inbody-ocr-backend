package usecase

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/internal/domain/service"
	"inbody-ocr-backend/internal/usecase/response"
	jptime "inbody-ocr-backend/pkg/jp_time"

	"github.com/uptrace/bun"
)

type measurementDateUsecase struct {
	repo                      repository.MeasurementDateRepository
	organizationRepo          repository.OrganizationRepository
	userMeasurementStatusRepo repository.UserMeasurementStatusRepository
	ulidService               service.ULIDService
}

func NewMeasurementDateUsecase(repo repository.MeasurementDateRepository, organizationRepo repository.OrganizationRepository, userMeasurementStatusRepo repository.UserMeasurementStatusRepository, ulidService service.ULIDService) MeasurementDateUsecase {
	return &measurementDateUsecase{
		repo:                      repo,
		organizationRepo:          organizationRepo,
		userMeasurementStatusRepo: userMeasurementStatusRepo,
		ulidService:               ulidService,
	}
}

func (uc *measurementDateUsecase) GetMeasurementDate(orgID string) (*response.GetMeasurementDateResponse, error) {
	dates, err := uc.repo.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	return response.NewGetMeasurementDateResponse(dates)
}

func (uc *measurementDateUsecase) CreateMeasurementDate(user *entity.User, dateStr string) (*response.CreateMeasurementDateResponse, error) {
	tx, err := uc.repo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			if rerr := tx.Rollback(); rerr != nil {
				// ロールバックに失敗した場合
				fmt.Printf("failed to rollback transaction: %v\n", rerr)
			}
			panic(p) // パニックが発生した場合はロールバックして再スロー
		} else if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				// ロールバックに失敗した場合
				fmt.Printf("failed to rollback transaction: %v\n", rerr)
			}
		}
	}()

	measurementDate, err := jptime.ParseDate(dateStr)
	if err != nil {
		return nil, err
	}
	measurementDateID := uc.ulidService.GenerateULID()
	// 測定日を保存
	date := entity.MeasurementDate{
		ID:             measurementDateID,
		OrganizationID: user.OrganizationID,
		Date:           measurementDate,
	}

	err = uc.repo.CreateMeasurementDateWithTx(tx, date)
	if err != nil {
		return nil, err
	}

	// 測定ステータスを作成
	count, err := uc.repo.CountByOrganizationID(user.OrganizationID)
	if err != nil {
		return nil, err
	}

	if count == 1 {
		err = uc.createMeasurementStatusesWithTx(tx, measurementDateID, user.OrganizationID)
		if err != nil {
			return nil, err
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return response.NewCreateMeasurementDateResponse(date)
}

// トランザクション内で測定ステータスを作成するヘルパー関数
func (uc *measurementDateUsecase) createMeasurementStatusesWithTx(tx bun.Tx, measurementDateID, organizationID string) error {
	members, err := uc.organizationRepo.GetMember(organizationID)
	if err != nil {
		return err
	}

	for _, member := range members {
		status := entity.UserMeasurementStatus{
			ID:                uc.ulidService.GenerateULID(),
			UserID:            member.ID,
			MeasurementDateID: measurementDateID,
			HasRegistered:     false,
		}

		err = uc.userMeasurementStatusRepo.CreateUserMeasurementStatusWithTx(tx, status)
		if err != nil {
			return err
		}
	}

	return nil
}
