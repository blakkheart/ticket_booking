package container

import (
	"ticket-booking/domain/service"
	"ticket-booking/models/database"
	"ticket-booking/repository/postgres"
)

type ContainerServiceStruct struct {
	AccountService *service.Service
}

var ContainerService *ContainerServiceStruct = &ContainerServiceStruct{}

func InitContainerService(dbConn *database.DBStruct) {

	// init repos
	accountRepo := &postgres.BaseRepository{DB: dbConn}

	// init services
	accountService := service.Service{Repo: accountRepo}

	ContainerService.AccountService = &accountService
}
