package service

type HealthCheckResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

//go:generate mockery --name HealthCheck --filename healthcheck_service.go
type HealthCheck interface {
	Check() (HealthCheckResponse, error)
}

type healthCheckService struct {
	serviceName string
	instanceID  string
}

func NewHealthCheck(serviceName, instanceID string) HealthCheck {
	return &healthCheckService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (s *healthCheckService) Check() (HealthCheckResponse, error) {
	return HealthCheckResponse{
		Message:     "OK",
		ServiceName: s.serviceName,
		InstanceID:  s.instanceID,
	}, nil
}
