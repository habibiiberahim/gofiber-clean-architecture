package repositorys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/gofiber-clean-architecture/entities"
	"github.com/habibiiberahim/gofiber-clean-architecture/pkg"
	"github.com/habibiiberahim/gofiber-clean-architecture/schemas"
	"gorm.io/gorm"
)

type RepositoryLogin interface{
	LoginRepository(input *schemas.SchemaAuth)(*entities.User, schemas.SchemaDatabaseError)
}
type repositoryLogin struct{
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repositoryLogin {
	return &repositoryLogin{
		db: db,
	}
}

func (r *repositoryLogin)LoginRepository(input *schemas.SchemaAuth)(*entities.User, schemas.SchemaDatabaseError)  {
	var user entities.User
	db := r.db.Model(&user)
	errorCode :=  make (chan schemas.SchemaDatabaseError, 1)

	user.Email = input.Email
	user.Password = input.Password

	checkUserAccount := db.Debug().First(&user,"email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_01",
		}
		return &user, <-errorCode
	}

	if !user.Active{
		errorCode <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_02",
		}
		return &user, <-errorCode
	}	

	comparePassword := pkg.ComparePassword(input.Password, user.Password )
	
	if comparePassword != nil {
		errorCode <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_03",
		}
		return &user, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}
	return &user, <-errorCode
}