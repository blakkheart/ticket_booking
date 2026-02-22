package config

import (
	"ticket-booking/domain/service"
	db "ticket-booking/repository"
)

type СontainerServiceStruct struct {
	Service *service.Service
}

var ContainerService СontainerServiceStruct = CreateContainerService()

func CreateContainerService() СontainerServiceStruct {
	// database := db.Connect()

	repo := &db.BaseRepository{}

	service := service.Service{Repo: repo}

	return СontainerServiceStruct{Service: &service}
}
