package dbutil

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Schema          string        `yaml:"schema"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Name            string        `yaml:"name"`
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time" default:"1h"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" default:"30m"`
	MaxOpenConns    int           `yaml:"max_open_conns" default:"4"`
	SSL             bool          `yaml:"ssl" default:"false"`
}

//func New(cfg Config, options ...option) (*DB, error) {
//	db := &DB{cfg: &cfg}
//	for _, option := range options {
//		if err := option(db); err != nil {
//			return nil, errors.Wrap(err, "apply option")
//		}
//	}
//	return db, nil
//}

func New(cfg Config) (*DB, error) {
	db := &DB{cfg: &cfg}
	return db, nil
}

type DB struct {
	pool     *pgxpool.Pool
	replicas []*DB
	tx       pgx.Tx
	cfg      *Config
	tracer   pgx.QueryTracer
}

func (conf *Config) String() string {
	format := "postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s"
	if conf.SSL {
		format = "postgres://%s:%s@%s:%d/%s?search_path=%s"
	}
	return fmt.Sprintf(format, conf.User, conf.Password, conf.Host, conf.Port, conf.Name, conf.Schema)
}

func (db *DB) Replica() *DB {
	if db.tx != nil || db.replicas == nil || len(db.replicas) == 0 {
		return db
	}
	return db.replicas[rand.Intn(len(db.replicas))]
}

func (db *DB) Tx(ctx context.Context) (*DB, error) {
	if db.tx != nil {
		return nil, errors.New("tx already started")
	}
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &DB{tx: tx}, nil
}

func (db *DB) Start(ctx context.Context) error {
	if db.tx != nil {
		return errors.New("cant start using tx")
	}
	var err error
	cfg, err := pgxpool.ParseConfig(db.cfg.String())
	if err != nil {
		return errors.Wrap(err, "pgxpool parse config")
	}
	cfg.MaxConnLifetime = db.cfg.ConnMaxLifeTime
	cfg.MaxConnIdleTime = db.cfg.ConnMaxIdleTime
	cfg.MaxConns = int32(db.cfg.MaxOpenConns)
	if db.tracer != nil {
		cfg.ConnConfig.Tracer = db.tracer
	}
	db.pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "pgxpool new with config")
	}
	if err := db.Ping(ctx); err != nil {
		return errors.Wrap(err, "db ping")
	}
	return nil
}

func (db *DB) Stop(ctx context.Context) error {
	if db.tx != nil {
		return errors.New("cant stop tx, use commit or rollback instead")
	}
	db.pool.Close()
	if db.replicas != nil {
		for _, replica := range db.replicas {
			_ = replica.Stop(ctx)
		}
	}
	return nil
}

func (db *DB) Rollback(ctx context.Context) error {
	if db.tx == nil {
		return errors.New("cant rollback on non tx")
	}
	return db.tx.Rollback(ctx)
}

func (db *DB) Commit(ctx context.Context) error {
	if db.tx == nil {
		return errors.New("cant commit on non tx")
	}
	return db.tx.Commit(ctx)
}

func (db *DB) Ping(ctx context.Context) error {
	if db.tx != nil {
		return errors.New("cant ping db using tx")
	}
	var err error
	if err = db.pool.Ping(ctx); err != nil {
		return err
	}
	if db.replicas != nil {
		for _, replica := range db.replicas {
			if err = replica.Ping(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if db.tx != nil {
		return db.tx.Query(ctx, query, args...)
	}
	return db.pool.Query(ctx, query, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if db.tx != nil {
		return db.tx.QueryRow(ctx, sql, args...)
	}
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *DB) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	if db.tx != nil {
		return db.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
	}
	return db.pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (db *DB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	if db.tx != nil {
		return db.tx.SendBatch(ctx, b)
	}
	return db.pool.SendBatch(ctx, b)
}

func (db *DB) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	if db.tx != nil {
		return db.tx.Exec(ctx, sql, arguments...)
	}
	return db.pool.Exec(ctx, sql, arguments...)
}
