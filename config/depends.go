package config

import (
	"ticket-booking/domain/service"
	db "ticket-booking/repository"
)

type containerServiceStruct struct {
	Service *service.Service
}

var ContainerService containerServiceStruct = CreateContainerService()

func CreateContainerService() containerServiceStruct {
	// database := db.Connect()

	repo := &db.BaseRepository{}

	service := service.Service{Repo: repo}

	return containerServiceStruct{Service: &service}
}
