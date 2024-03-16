package lib

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	C "github.com/gabriel-tama/be-week-1/config"
)

var PgPool *pgxpool.Pool

func GetPostgresURL() string {
	env, err := C.Get()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPassword
	dbName := env.DBName

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass,
		dbHost, dbPort, dbName)
}

func Init(ctx context.Context) error {
	var err error

	PgPool, err = pgxpool.New(context.Background(), GetPostgresURL())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func PGTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := PgPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func Close(ctx context.Context) {
	PgPool.Close()
}
