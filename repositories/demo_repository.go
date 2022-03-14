package repositories

import (
	"webapi/db"
	"webapi/models/domains"
)

type IDemoRepository interface {
	IBaseRepository
}

type DemoRepository struct {
	*BaseRepository
}

func NewDemoRepository(db *db.DB) IDemoRepository {
	return &DemoRepository{
		&BaseRepository{
			db: db.Table(domains.DemoTableName),
		},
	}
}
