package database

import (
	"cards-binder/internal/catalog"
	"database/sql"
	"fmt"
)

type CardDataBase interface {
	NewCard(card *catalog.Card) (*catalog.Card, error)
	GetCard(id string) (*catalog.Card, error)
}

func NewCardDataBase() CardDataBase {
	if isNative() {
		return newNativeCardDataBase()
	} else {
		return newGormCardDataBase()
	}
}

func isNative() bool {
	return false
}

func NewSqlClient(drv string, datasource string) *sql.DB {
	db, err := sql.Open(drv, datasource)

	if err != nil {
		_ = fmt.Errorf("Cannot create db tentat: %s", err.Error())
		panic("No se pudo establecer conección con la base de datos")
	}

	return db
}
