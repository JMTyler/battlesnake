package db

import (
	"context"
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/go-pg/pg/v9"
	"net/url"
	"strings"
)

var DB *pg.DB

var queue = make(chan func())

func InitDatabase() *pg.DB {
	if DB != nil {
		return DB
	}

	databaseUrl := config.Get("database_url", "")
	if databaseUrl == "" {
		// TODO: throw error, need database
		panic("need database")
	}

	// TODO: Might actually be able to use pg.Options.ParseURL()
	auth, err := url.Parse(databaseUrl)
	if err != nil {
		panic(err)
	}

	if auth.Scheme != "postgres" {
		// TODO: throw error, DB must be postgres
		panic("gots ta be PG")
	}

	password, _ := auth.User.Password()
	database := strings.TrimPrefix(auth.Path, "/")

	connOptions := &pg.Options{
		User:            auth.User.Username(),
		Password:        password,
		Addr:            auth.Host,
		Database:        database,
		ApplicationName: "BattlesnakeServer",
	}

	DB = pg.Connect(connOptions)

	if config.GetBool("debug") {
		DB.AddQueryHook(dbLogger{})
		//		pg.SetLogger(new(dbLogger))
	}

	if err := migrate(); err != nil {
		panic(err)
	}

	return DB
}

func CloseDatabase() {
	if DB != nil {
		fmt.Println("Closing database connection...")
		DB.Close()
	}
}

func WatchQueue() {
	for operation := range queue {
		operation()
	}
}

func CloseQueue() {
	close(queue)

	// Clear out the queue before moving on.
	for operation := range queue {
		operation()
	}

	fmt.Println("Database queue closed.")
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	sql, _ := q.FormattedQuery()
	fmt.Println("[SQL]", sql)
	return ctx, nil
}

func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	return nil
}
