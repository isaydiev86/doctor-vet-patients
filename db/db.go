package db

import "github.com/jmoiron/sqlx"

type DB struct {
	*sqlx.DB
}

func New(cfg Config) (*DB, error) {
	/*db, err := sqlx.Connect("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
	/**/
	return &DB{}, nil
}
