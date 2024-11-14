package response

import "inbody-ocr-backend/internal/domain/entity"

type GetMeasurementDateResponse struct {
	MeasurementDates []MeasurementDateResponse `json:"measurement_dates"`
}

func NewGetMeasurementDateResponse(measurementDates []entity.MeasurementDate) (*GetMeasurementDateResponse, error) {
	return &GetMeasurementDateResponse{
		MeasurementDates: NewMeasurementDateResponseList(measurementDates),
	}, nil
}
