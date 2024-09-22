//go:build wireinject
// +build wireinject

package main

import (
	db "golang_grafana_struct/infra"
	"golang_grafana_struct/passage"

	passagestore "golang_grafana_struct/passage/database"
	passageservice "golang_grafana_struct/passage/service"

	"github.com/google/wire"
)

var wireBasicSet = wire.NewSet(
	db.ProvideDb,
	passagestore.ProvidePassageStore,
	passageservice.ProvidePassageService,
)

func InitializePassageService() (passage.PassageService, error) {
	wire.Build(wireBasicSet)
	return &passageservice.PassageServiceImpl{}, nil
}
