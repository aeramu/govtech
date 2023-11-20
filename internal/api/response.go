package api

type MutationResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
