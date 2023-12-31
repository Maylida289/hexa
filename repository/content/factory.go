package content

import (
	"hexa/util"

	"hexa/business/content"
	// "google.golang.org/api/content/v2"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
	contentRepo := NewPostgresRepository(dbCon.Postgres)
	return contentRepo
}
