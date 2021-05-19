package models

// HealthCheckResponse...
type ApiResponse struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}
