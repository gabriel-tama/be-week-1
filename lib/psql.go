package lib

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"

	C "github.com/gabriel-tama/be-week-1/config"
)

var PostgresConn *pgx.Conn

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
	PostgresConn, err = pgx.Connect(ctx, GetPostgresURL())
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	err = PostgresConn.Ping(ctx)
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

func PGTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := PostgresConn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func Close(ctx context.Context) {
	PostgresConn.Close(ctx)
}
