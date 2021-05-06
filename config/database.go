package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ternaknesia/qolbu/exception"
)

func configDatabase(configuration Config, target string) (string, string, string, string) {
	dbHost := configuration.Get("DB_HOST" + target)
	dbUser := configuration.Get("DB_USERNAME" + target)
	dbPass := configuration.Get("DB_PASSWORD" + target)
	dbName := configuration.Get("DB_DATABASE" + target)
	return dbHost, dbUser, dbPass, dbName
}

func createDSN(configuration Config, target string) string {
	dbHost, dbUser, dbPass, dbName := configDatabase(configuration, target)
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
}

func openDatabase(ctx context.Context, configuration Config, target string) *sql.DB {
	db, err := sql.Open("mysql", createDSN(configuration, target))
	exception.PanicIfNeeded(err)

	err = db.PingContext(ctx)
	exception.PanicIfNeeded(err)

	return db
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 300*time.Second)
}

func CreateDatabase(configuration Config, target string) (*sql.DB, context.Context, context.CancelFunc) {
	ctx, cancel := createContext()
	defer cancel()
	return openDatabase(ctx, configuration, target), ctx, cancel
}
