package database

import "cards-binder/internal/catalog"

type nativeCardDataBase struct{}

func newNativeCardDataBase() *nativeCardDataBase {
	return &nativeCardDataBase{}
}

func (n *nativeCardDataBase) NewCard(card *catalog.Card) (*catalog.Card, error) {
	return nil, nil
}

func (n *nativeCardDataBase) GetCard(id string) (*catalog.Card, error) {
	return nil, nil
}
