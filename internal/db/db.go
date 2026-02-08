package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type DSN struct {
	database string
	username string
	password string
	host     string
	port     int
	schema   string
}

type DSNOpt func(*DSN)

func NewDSN(opts ...DSNOpt) *DSN {
	dsn := &DSN{
		host:     "localhost",
		port:     5432,
		database: "postgres",
		username: "postgres",
		password: "password",
		schema:   "public",
	}
	for _, o := range opts {
		o(dsn)
	}

	return dsn
}

func WithHost(host string) DSNOpt {
	return func(d *DSN) {
		if host != "" {
			d.host = host
		}
	}
}

func WithDatabase(database string) DSNOpt {
	return func(d *DSN) {
		if database != "" {
			d.database = database
		}
	}
}

func WithUsername(username string) DSNOpt {
	return func(d *DSN) {
		if username != "" {
			d.username = username
		}
	}
}

func WithPassword(password string) DSNOpt {
	return func(d *DSN) {
		if password != "" {
			d.password = password
		}
	}
}

func WithPort(port int) DSNOpt {
	return func(d *DSN) {
		if port > 0 {
			d.port = port
		}
	}
}

func (d *DSN) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.username, d.password, d.host, d.port, d.database)
}

func Connect(ctx context.Context, opts ...DSNOpt) (*pgx.Conn, error) {
	dsn := NewDSN(opts...)
	conn, err := pgx.Connect(ctx, dsn.ConnString())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Query(ctx context.Context, query string, connOpts ...DSNOpt) (*sql.Rows, error) {
	conn, err := Connect(ctx, connOpts...)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	db := stdlib.OpenDB(*conn.Config())
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
