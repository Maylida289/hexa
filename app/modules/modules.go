//untuk registrasi modul atau implementasi dependendy injection nya
package modules

import (
	"hexa/api"
	"hexa/util"

	contentV1Controller "hexa/api/v1/content"
	contentV2Controller "hexa/api/v2/content"
	contentservice "hexa/business/content"
	contentRepo "hexa/repository/content"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentservice.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	contentV2PermitController := contentV2Controller.NewController(contentPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		ContentV2Controller: contentV2PermitController,
	}

	return controllers
}
