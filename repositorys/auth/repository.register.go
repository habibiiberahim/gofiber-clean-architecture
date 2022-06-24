package repositorys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/entities"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
	"gorm.io/gorm"
)

type RepositoryRegister interface {
	// input from schemas.schemaAuth and return entities.EntityUser and schemas.SchemaDatabaseError
	RegisterRepository(input *schemas.SchemaAuth)(*entities.EntityUser, schemas.SchemaDatabaseError)	
}

type repositoryRegister struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repositoryRegister  {
	return &repositoryRegister{
		db: db,
	}
}

func (r *repositoryRegister)RegisterRepository(input *schemas.SchemaAuth) (*entities.EntityUser, schemas.SchemaDatabaseError)  {
	var user entities.EntityUser
	db := r.db.Model(&user)	
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)
	if checkUserAccount.RowsAffected>0 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: fiber.StatusConflict,
			Type: "error_01",
		}
		return &user, <- errorCode
	}
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Password = input.Password

	addNewuser := db.Debug().Create(&user).Commit()

	if addNewuser.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_02",
		}
		return &user, <- errorCode
	}
	
	return &user, <- errorCode
}
