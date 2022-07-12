package item

import (
	"errors"
	"fmt"

	"github.com/aaalik/api-alik/internal/api/constants"
	"github.com/aaalik/api-alik/internal/api/contracts"
	"github.com/aaalik/api-alik/internal/api/entities"
	"github.com/aaalik/api-alik/pkg/alog"
)

type Service struct {
	repo contracts.ItemRepository
}

func Init(app *contracts.App) (svc contracts.ItemService) {
	r := initRepository(app.Datasources.WriterDB, app.Datasources.ReaderDB, app.Config[constants.AesKey])

	svc = &Service{
		repo: r,
	}

	return
}

func (s *Service) GetItems() (items []entities.Item, err error) {
	items, err = s.repo.GetItems()
	return
}

func (s *Service) FindItemById(id int) (item entities.Item, err error) {
	item, err = s.repo.FindItemById(id)
	return
}

func (s *Service) InsertItem(item entities.Item) (entities.Item, error) {
	item, err := s.repo.InsertItem(item)

	if err != nil {
		alog.Logger.Error(errors.New(fmt.Sprintf("fail to insert item, err: %v", err)))
		return entities.Item{}, err
	}

	return item, err
}

func (s *Service) UpdateItem(id int, item map[string]interface{}) (entities.Item, error) {
	resItem, err := s.repo.UpdateItem(id, item)
	return resItem, err
}

func (s *Service) DeleteItem(id int) (err error) {
	err = s.repo.DeleteItem(id)
	return
}
