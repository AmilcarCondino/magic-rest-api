package database

import "cards-binder/internal/catalog"

type gormCardDataBase struct{}

func newGormCardDataBase() *gormCardDataBase {
	return &gormCardDataBase{}
}

func (n *gormCardDataBase) NewCard(card *catalog.Card) (*catalog.Card, error) {
	return nil, nil
}

func (n *gormCardDataBase) GetCard(id string) (*catalog.Card, error) {
	return nil, nil
}
