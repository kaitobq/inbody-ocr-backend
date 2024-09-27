package response

type SaveImageDataResponse struct {
	Message string `json:"message"`
}

func NewSaveImageDataResponse() (*SaveImageDataResponse, error) {
	return &SaveImageDataResponse{
		Message: "Data saved successfully",
	}, nil
}
