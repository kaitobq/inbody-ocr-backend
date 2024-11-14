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

type CreateMeasurementDateResponse struct {
	MeasurementDate MeasurementDateResponse `json:"measurement_date"`
}

func NewCreateMeasurementDateResponse(measurementDate entity.MeasurementDate) (*CreateMeasurementDateResponse, error) {
	return &CreateMeasurementDateResponse{
		MeasurementDate: NewMeasurementDateResponse(measurementDate),
	}, nil
}
