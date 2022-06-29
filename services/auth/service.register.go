package services

import (
	"github.com/habibiiberahim/gofiber-clean-architecture/entities"
	repositorys "github.com/habibiiberahim/gofiber-clean-architecture/repositorys/auth"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
)

//create interface
type ServiceRegister interface{
	RegisterService(input *schemas.SchemaAuth)(*entities.User, schemas.SchemaDatabaseError)
}

//create struct
type serviceRegister struct {
	repository repositorys.RepositoryRegister
}

//create init func
func NewServiceRegister(repository repositorys.RepositoryRegister) *serviceRegister  {
	return &serviceRegister{
		repository: repository,
	}
}

//all function need 
func (s *serviceRegister)RegisterService(input *schemas.SchemaAuth) (*entities.User,schemas.SchemaDatabaseError) {
	var schema schemas.SchemaAuth
	schema.Fullname = input.Fullname
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.RegisterRepository(&schema)
	return res, err
}
