package service

// SkyEyeService ...
var SkyEyeService *Service

// Init ...
func Init(config *ServiceConfig) {

	SkyEyeService = NewService(config)

	return
}
