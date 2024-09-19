package response

import "time"

type OrganizationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrganizationResponse struct {
	Message   string `json:"message"`
	Organization OrganizationResponse `json:"organization"`
}

func NewCreateOrganizationResponse(id string, name string, createdAt time.Time, updatedAt time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
		Message: "Organization created successfully",
		Organization: OrganizationResponse{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}
