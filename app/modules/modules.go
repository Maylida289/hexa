//untuk registrasi modul atau implementasi dependendy injection nya
package modules

import (
	"hexa/api"
	"hexa/util"

	// "github.com/swaggo/swag/example/basic/api"
	contentController "hexa/api/v1/content"
	contentservice "hexa/business/content"
	contentRepo "hexa/repository/content"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentservice.NewService(contentPermitRepository)
	contentPermitController := contentController.NewController(contentPermitService)

	controllers := api.Controller{
		ContentController: contentPermitController,
	}

	return controllers
}
