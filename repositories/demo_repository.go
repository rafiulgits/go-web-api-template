package repositories

import (
	"webapi/db"
	"webapi/models/domains"
)

type IDemoRepository interface {
	IBaseRepository[int, domains.Demo]
}

type DemoRepository struct {
	*BaseRepository[int, domains.Demo]
}

func NewDemoRepository(db *db.DB) IDemoRepository {
	return &DemoRepository{
		&BaseRepository[int, domains.Demo]{
			db:        db.DB,
			tableName: domains.DemoTableName,
		},
	}
}
