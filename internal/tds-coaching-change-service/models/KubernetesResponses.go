package models

// HealthResponse represents the response for a health check.
// swagger:response HealthResponse
type HealthResponse struct {
	Status string `json:"status"`
}

// ReadinessResponse represents the response for a readiness check.
// swagger:response ReadinessResponse
type ReadinessResponse struct {
	Status string `json:"status"`
}

// StartupResponse represents the response for a startup check.
// swagger:response StartupResponse
type StartupResponse struct {
	Status string `json:"status"`
}
