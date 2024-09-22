package database

import (
	db "golang_grafana_struct/infra"
	"golang_grafana_struct/passage"
)

type PassageStore struct {
	store db.DB
}

// PassageStore implements the Store interface
var _ passage.Store = (*PassageStore)(nil)

func ProvidePassageStore(db db.DB) (passage.Store, error) {
	s := &PassageStore{store: db}

	return s, nil
}

func (p *PassageStore) GetPassageById(query passage.GetPassageQuery) int {
	p.store.Query("SELECT * FROM EXAMPLE")
	return 1
}
func (p *PassageStore) SavePassage(cmd passage.SavePassageCommand) int {
	p.store.Query("INSERT x (y,z) values (a,b)")
	return 1
}
