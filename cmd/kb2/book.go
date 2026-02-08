package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jansdhillon/kb2/.gen/postgres/public/model"
	"github.com/jansdhillon/kb2/.gen/postgres/public/table"
	"github.com/urfave/cli/v3"
)

var getBooksCmd = &cli.Command{
	Name:  "get",
	Usage: "Get books from the database",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		db, ok := ctx.Value(dbClientKey).(*sql.DB)
		if !ok || db == nil {
			return fmt.Errorf("DB client not initialized")
		}

		stmt := SELECT(
			table.Books.AllColumns,
		).FROM(table.Books).ORDER_BY(table.Books.CreatedAt.ASC())

		var books []model.Books
		err := stmt.QueryContext(ctx, db, &books)
		if err != nil {
			return fmt.Errorf("failed to query db: %s", err)
		}

		log.Printf("books: %v", books)

		return nil

	},
}
