package main

import (
	"github.com/aaalik/api-alik/pkg/alog"
	"github.com/joho/godotenv"
)

func initConfig() {
	// Used if using .env file
	env, err := godotenv.Read()
	if err != nil {
		alog.Logger.Fatalf(err.Error())
	} else {
		alog.Logger.Printf("Initalizing Config: Pass")
	}

	app.Config = env
}
