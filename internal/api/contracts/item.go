package contracts

import "github.com/aaalik/api-alik/internal/api/entities"

type ItemRepository interface {
	GetItems() (items []entities.Item, err error)
	FindItemById(id int) (item entities.Item, err error)
	InsertItem(item entities.Item) (entities.Item, error)
	UpdateItem(id int, item map[string]interface{}) (entities.Item, error)
	DeleteItem(id int) (err error)
}

type ItemService interface {
	GetItems() (items []entities.Item, err error)
	FindItemById(id int) (item entities.Item, err error)
	InsertItem(item entities.Item) (entities.Item, error)
	UpdateItem(id int, item map[string]interface{}) (entities.Item, error)
	DeleteItem(id int) (err error)
}
