package storage

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStorage{db},
		Users: &UserStorage{db},
	}
}
