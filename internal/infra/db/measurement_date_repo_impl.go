package db

import (
	"inbody-ocr-backend/internal/domain/entity"
	"inbody-ocr-backend/internal/domain/repository"
	"inbody-ocr-backend/pkg/database"
	jptime "inbody-ocr-backend/pkg/jp_time"
)

type measurementDateRepository struct {
	db database.DB
}

func NewMeasurementDateRepository(db *database.DB) repository.MeasurementDateRepository {
	return &measurementDateRepository{
		db: *db,
	}
}

func (r *measurementDateRepository) FindByOrganizationID(orgID string) ([]entity.MeasurementDate, error) {
	query := `SELECT id, organization_id, DATE_FORMAT(date, '%Y-%m-%d') as date, created_at, updated_at FROM measurement_date WHERE organization_id = ?`

	var dates []entity.MeasurementDate

	rows, err := r.db.Query(query, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record entity.MeasurementDate
		var dateStr string
		var createdAt, updatedAt string

		// DATEフィールドを文字列として取得
		err := rows.Scan(&record.ID, &record.OrganizationID, &dateStr, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		// 文字列をtime.Timeに変換
		record.Date, err = jptime.ParseDate(dateStr)
		if err != nil {
			return nil, err
		}

		// 作成日と更新日のパース
		record.CreatedAt, err = jptime.ParseDateTime(createdAt)
		if err != nil {
			return nil, err
		}
		record.UpdatedAt, err = jptime.ParseDateTime(updatedAt)
		if err != nil {
			return nil, err
		}

		dates = append(dates, record)
	}

	return dates, nil
}
