package routers

import (
	"github.com/aaalik/api-alik/internal/api/contracts"
	"github.com/aaalik/api-alik/internal/api/handlers"
	"github.com/kataras/iris/v12"
)

func Init(app *contracts.App, crs iris.Handler) {

	app.Iris.Get("/", func(c iris.Context) {
		handlers.HttpSuccess(c, nil, "")
	})

	party := app.Iris.Party("/v1", crs)
	{
		party.Get("/validate", func(c iris.Context) {
			handlers.HttpSuccess(c, nil, "")
		})

		party.Get("/items", handlers.GetItems)
		party.Get("/items/{id}", handlers.FindItemById)
		party.Put("/items/{id}", handlers.UpdateItem)
		party.Post("/items", handlers.InsertItem)
		party.Delete("/items/{id}", handlers.DeleteItem)
	}
}
