package dto

type (
	AuthorizationRequest struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}
)
