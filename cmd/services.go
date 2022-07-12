package main

import (
	"github.com/aaalik/api-alik/internal/api/contracts"
	"github.com/aaalik/api-alik/internal/api/services/item"
	"github.com/aaalik/api-alik/pkg/alog"
)

func initServices() {
	app.Services = &contracts.Services{
		Item: item.Init(app),
	}

	alog.Logger.Printf("Initializing Services: Pass")
}
