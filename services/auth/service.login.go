package services

import (
	"github.com/habibiiberahim/gofiber-clean-architecture/entities"
	repositorys "github.com/habibiiberahim/gofiber-clean-architecture/repositorys/auth"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
)

type ServiceLogin interface {
	LoginService(input *schemas.SchemaAuth)(*entities.User, schemas.SchemaDatabaseError)
}

type serviceLogin struct {
	repository repositorys.RepositoryLogin
}

func NewServiceLogin(repository repositorys.RepositoryLogin) *serviceLogin {
	return &serviceLogin{
		repository: repository,
	}
}

func (s *serviceLogin) LoginService(input *schemas.SchemaAuth) (*entities.User, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.LoginRepository(&schema)

	return res, err
}