package handlers

import (
	"strconv"

	"github.com/aaalik/api-alik/internal/api/entities"
	"github.com/aaalik/api-alik/pkg/ahttp"
	"github.com/kataras/iris/v12"
)

func GetItems(c iris.Context) {
	// get item
	items, err := app.Services.Item.GetItems()
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	HttpSuccess(c, items, "asdgsadfgs")
}

func FindItemById(c iris.Context) {
	// get item
	strId := c.Params().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	item, err := app.Services.Item.FindItemById(id)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	HttpSuccess(c, item, "")
}

func InsertItem(c iris.Context) {
	// get item
	var item entities.Item
	err := c.ReadJSON(&item)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	item, err = app.Services.Item.InsertItem(item)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	HttpSuccess(c, item, "")
}

func UpdateItem(c iris.Context) {
	var err error
	// get item
	strId := c.Params().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	var item map[string]interface{}
	err = c.ReadJSON(&item)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	var resItem entities.Item
	resItem, err = app.Services.Item.UpdateItem(id, item)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	HttpSuccess(c, resItem, "")
}

func DeleteItem(c iris.Context) {
	// get item
	strId := c.Params().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	err = app.Services.Item.DeleteItem(id)
	if err != nil {
		HttpError(c, ahttp.ErrInternalServer, err)
		return
	}

	HttpSuccess(c, "", "")
}
