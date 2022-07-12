package handlers

import (
	. "github.com/aaalik/api-alik/internal/api/contracts"
)

var app *App

func Init(a *App) {
	app = a
}
