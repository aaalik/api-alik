package item

import (
	"errors"
	"fmt"
	"time"

	"github.com/aaalik/api-alik/internal/api/contracts"
	"github.com/aaalik/api-alik/internal/api/datasources"
	"github.com/aaalik/api-alik/internal/api/entities"
	"github.com/aaalik/api-alik/pkg/alog"
	"github.com/aaalik/api-alik/pkg/autils"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	dbWriter *sqlx.DB
	dbReader *sqlx.DB
	stmt     Statement
}

type Statement struct {
	getItems     *sqlx.Stmt
	findItemById *sqlx.Stmt
	insertItem   *sqlx.NamedStmt
	updateItem   *sqlx.NamedStmt
	deleteItem   *sqlx.NamedStmt
}

func initRepository(dbWriter *sqlx.DB, dbReader *sqlx.DB, cryptKey string) contracts.ItemRepository {

	stmts := Statement{
		getItems:     datasources.Prepare(dbReader, getItems),
		findItemById: datasources.Prepare(dbReader, findItemById),
		insertItem:   datasources.PrepareNamed(dbWriter, insertItem),
		updateItem:   datasources.PrepareNamed(dbWriter, updateItem),
		deleteItem:   datasources.PrepareNamed(dbWriter, deleteItem),
	}

	r := Repository{
		dbWriter: dbWriter,
		dbReader: dbReader,
		stmt:     stmts,
	}

	return &r
}

func (r Repository) GetItems() (items []entities.Item, err error) {
	err = r.stmt.getItems.Select(&items)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to get items, err: %v", err)))
		return
	}

	return
}

func (r Repository) FindItemById(id int) (item entities.Item, err error) {
	row := r.stmt.findItemById.QueryRow(id)

	err = row.Scan(
		&item.Id,
		&item.Name,
		&item.Price,
		&item.Description,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to find item, err: %v", err)))
		return
	}

	return
}

func (r Repository) InsertItem(item entities.Item) (entities.Item, error) {
	args := map[string]interface{}{
		"name":        item.Name,
		"price":       item.Price,
		"description": item.Description,
	}

	row, err := r.stmt.insertItem.Exec(args)

	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to insert item, err: %v", err)))
		return entities.Item{}, err
	}

	lastInsertedID, err := row.LastInsertId()

	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to insert item, err: %v", err)))
		return entities.Item{}, err
	}

	item.Id = int(lastInsertedID)
	item.CreatedAt = time.Now().Format(autils.DATETIME_LAYOUT)

	return item, err
}

func (r Repository) UpdateItem(id int, item map[string]interface{}) (entities.Item, error) {
	var err error
	var resItem entities.Item

	row := r.stmt.findItemById.QueryRow(id)

	err = row.Scan(
		&resItem.Id,
		&resItem.Name,
		&resItem.Price,
		&resItem.Description,
		&resItem.CreatedAt,
		&resItem.UpdatedAt,
	)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to find item, err: %v", err)))
		return entities.Item{}, err
	}

	for i, v := range item {
		if i == "name" {
			resItem.Name = v.(string)
		}

		if i == "price" {
			resItem.Price = int(v.(float64))
		}

		if i == "description" {
			resItem.Description = v.(string)
		}
	}

	args := map[string]interface{}{
		"name":        resItem.Name,
		"price":       resItem.Price,
		"description": resItem.Description,
		"id":          id,
	}

	_, err = r.stmt.updateItem.Exec(args)
	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to update item, err: %v", err)))
		return entities.Item{}, err
	}

	return resItem, err
}

func (r Repository) DeleteItem(id int) (err error) {
	args := map[string]interface{}{
		"id": id,
	}

	_, err = r.stmt.deleteItem.Exec(args)

	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to delete item, err: %v", err)))
		return
	}

	return
}
