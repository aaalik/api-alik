package main

import (
	"net/http"
	"os"
	"time"

	"github.com/aaalik/api-alik/internal/api/constants"
	"github.com/aaalik/api-alik/internal/api/contracts"
	"github.com/aaalik/api-alik/internal/api/handlers"
	"github.com/aaalik/api-alik/internal/api/middlewares"
	"github.com/aaalik/api-alik/internal/api/routers"
	"github.com/aaalik/api-alik/pkg/alog"
	_ "github.com/aaalik/api-alik/pkg/tzinit"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

var app *contracts.App

func main() {
	os.Setenv("TZ", "Asia/Jakarta")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Screen"},
		AllowCredentials: true,
	})

	irisApp := iris.New()
	irisApp.Use(crs)
	irisApp.AllowMethods(iris.MethodOptions)

	app = &contracts.App{
		Iris: irisApp,
	}

	alog.Init()

	initConfig()
	initDatasource()
	initServices()

	middlewares.Init(app)
	handlers.Init(app)
	routers.Init(app, crs)

	srv := &http.Server{
		Addr:         ":" + app.Config[constants.AppPort],
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	irisApp.Run(iris.Server(srv), iris.WithOptimizations, iris.WithoutBodyConsumptionOnUnmarshal)
}
