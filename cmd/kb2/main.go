package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jansdhillon/kb2/internal/db"
	"github.com/urfave/cli/v3"
)

const (
	databaseFlag   = "database"
	dbHostFlag     = "host"
	dbUsernameFlag = "username"
	dbPasswordFlag = "password"
	dBPortFlag     = "port"
)

const dbClientKey = "kb2-db-client"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := &cli.Command{
		Name:    "kb2",
		Version: "v0.0.1",
		Usage:   "Root entry point for kb2",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    databaseFlag,
				Aliases: []string{"d"},
				Usage:   "The name of the database to connect to",
				Sources: cli.EnvVars("KB2_DB_NAME"),
			},
			&cli.StringFlag{
				Name:    dbHostFlag,
				Aliases: []string{"h"},
				Usage:   "The host of the database to connect to",
				Sources: cli.EnvVars("KB2_DB_HOST"),
			},
			&cli.StringFlag{
				Name:    dbUsernameFlag,
				Aliases: []string{"u"},
				Usage:   "The username of the user to perform the query as",
				Sources: cli.EnvVars("KB2_DB_USERNAME"),
			},
			&cli.StringFlag{
				Name:    dbPasswordFlag,
				Usage:   "The password of the user to perform the query as",
				Sources: cli.EnvVars("KB2_DB_PASSWORD"),
			},
			&cli.StringFlag{
				Name:    dBPortFlag,
				Usage:   "The port of the database to connect to",
				Sources: cli.EnvVars("KB2_DB_PORT"),
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			database := cmd.String(databaseFlag)
			host := cmd.String(dbHostFlag)
			username := cmd.String(dbUsernameFlag)
			password := cmd.String(dbPasswordFlag)
			port := cmd.Int(dBPortFlag)

			if database == "" {
				return ctx, fmt.Errorf("database name must be provided")
			}
			if username == "" {
				return ctx, fmt.Errorf("database username must be provided")
			}
			if password == "" {
				return ctx, fmt.Errorf("database password must be provided")
			}

			if host == "" {
				host = "localhost"
			}

			if port == 0 {
				port = 5432
			}

			conn, err := db.Connect(ctx, db.WithDatabase(database), db.WithHost(host), db.WithUsername(username), db.WithPassword(password), db.WithPort(port))
			if err != nil {
				return ctx, fmt.Errorf("failed to connect to database: %s", err)
			}

			defer conn.Close(ctx)

			db := stdlib.OpenDB(*conn.Config())

			return context.WithValue(ctx, dbClientKey, db), nil

		},
		Commands: []*cli.Command{
			getBooksCmd,
		},
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
