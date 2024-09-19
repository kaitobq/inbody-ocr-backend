package response

type DeleteMembershipResponse struct {
	Message string `json:"message"`
}

func NewDeleteMembershipResponse() (*DeleteMembershipResponse, error) {
	return &DeleteMembershipResponse{
		Message: "Membership deleted successfully",
	}, nil
}
