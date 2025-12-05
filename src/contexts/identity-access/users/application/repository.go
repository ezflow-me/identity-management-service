package application

import (
	"os"

	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/domain"
	"github.com/ezflow-me/identity-management-service/src/contexts/identity-access/users/infrastructure/persistence/postgre"
)

var Repository domain.UserRepository

func init() {
	switch os.Getenv("DATABASE") {
	case "POSTGRE":
		Repository = postgre.NewPostgreRepository(os.Getenv("POSTGRE_URL"))
	default:
		panic("Database not configured")
	}
}
