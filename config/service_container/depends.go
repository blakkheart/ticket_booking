package container

import (
	"ticket-booking/domain/service"
	"ticket-booking/repository/postgres"
)

type ContainerServiceStruct struct {
	AccountService *service.Service
}

var ContainerService *ContainerServiceStruct = CreateContainerService()

func CreateContainerService() *ContainerServiceStruct {
	// database := db.Connect()

	accountRepo := &postgres.BaseRepository{}
	accountService := service.Service{Repo: accountRepo}

	return &ContainerServiceStruct{AccountService: &accountService}
}
