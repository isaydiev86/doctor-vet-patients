package db

import (
	"context"

	"doctor-vet-patients/pkg/dbutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type DB struct {
	*dbutil.DB
	cfg    dbutil.Config
	logger *zap.Logger
}

func New(cfg dbutil.Config, logger *zap.Logger) (*DB, error) {
	db, err := dbutil.New(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "dbutil new")
	}
	return &DB{DB: db, cfg: cfg, logger: logger}, nil
}

func (db *DB) Start(ctx context.Context) error { return db.DB.Start(ctx) }
func (db *DB) Stop(ctx context.Context) error  { return db.DB.Stop(ctx) }

func (db *DB) Tx(ctx context.Context, f func(any) error) error {
	tx, err := db.DB.Tx(ctx)
	if err != nil {
		return err
	}
	txDB := &DB{DB: tx, cfg: db.cfg}
	if err := f(txDB); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
