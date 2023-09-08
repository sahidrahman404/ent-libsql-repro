package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"todo"
	"todo/ent"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/libsql/libsql-client-go/libsql"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	dsn := "libsql://"

	db, err := sql.Open(
		"libsql",
		dsn,
	)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	drv := entsql.OpenDB(dialect.SQLite, db)
	client := ent.NewClient(ent.Driver(drv))

	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		// migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", err)
	}
	srv := handler.NewDefaultServer(todo.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on", cli.Addr)
	if err := http.ListenAndServe(cli.Addr, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
