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

func createDSN(configuration Config, target string) (string, string) {
	dbHost, dbUser, dbPass, dbName := configDatabase(configuration, target)
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", 
		dbUser, dbPass, dbHost, dbName),
		dbName
}

func openDatabase(ctx context.Context, configuration Config, target string) (*sql.DB, string) {
	dsn, dbName := createDSN(configuration, target)
	db, err := sql.Open("mysql", dsn)
	exception.PanicIfNeeded(err)

	err = db.PingContext(ctx)
	exception.PanicIfNeeded(err)

	return db, dbName
}

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 300*time.Second)
}

func CreateDatabase(configuration Config, target string) (*sql.DB, string) {
	ctx, cancel := CreateContext()
	defer cancel()
	return openDatabase(ctx, configuration, target)
}
