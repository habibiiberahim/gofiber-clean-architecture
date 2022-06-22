package services

import repositorys "github.com/habibiiberahim/gofiber-clean-architecture/repositorys/auth"

type ServicePing interface {
	PingService() string
}

type servicePing struct{
	repository repositorys.RepositoryPing
}

func NewServicePing(repository repositorys.RepositoryPing) *servicePing{
	return &servicePing{repository: repository}
}

func (s *servicePing)PingService()string  {
	res := s.repository.PingRepository()
	return res
}