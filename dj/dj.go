package dj

import (
	"go_restfull_api/database"
	"go_restfull_api/repositories"
	"go_restfull_api/services"
	"go_restfull_api/utils"
)

func InitDjContainer() {
	camRepo := repositories.CameraRepository{
		DB: database.GetMysqlDB(),
	}

	camService := services.CameraService{
		CameraRepository: &camRepo,
	}

	utils.Container.BindObject(&camRepo)
	utils.Container.BindObject(&camService)
}
