package middlewares

import "github.com/aaalik/api-alik/internal/api/contracts"

var app *contracts.App

func Init(a *contracts.App) {
	app = a
}
