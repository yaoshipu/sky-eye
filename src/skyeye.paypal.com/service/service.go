package service

// Service struct
type Service struct {
	config *ServiceConfig
}

// NewService constructor
func NewService(cfg *ServiceConfig) *Service {
	return &Service{config: cfg}
}
