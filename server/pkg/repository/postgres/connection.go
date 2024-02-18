package postgres

import (
	"context"
	"log"
	"time"

	config "github.com/AfsanehHabibi/neveshtedan/pkg/config/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitializeDBPool() error {
	var err error

	dbConfig, err := pgxpool.ParseConfig(config.PgConnStr)
	if err != nil {
		return err
	}

	dbConfig.MaxConns = int32(config.GetConfigValue("PG_MAX_CON"))
	dbConfig.MaxConnIdleTime = time.Duration(config.GetConfigValue("PG_MAX_CON_IDLE_TIME"))

	dbPool, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return err
	}

	return nil
}

func DB() *pgxpool.Pool {
	return dbPool
}

func Close() {
	dbPool.Close()
}

func Setup() error {
	retryCount := 10
	for i := 1; i <= retryCount; i++ {
		err := ConnectToPostgresql()
		if err == nil {
			break
		}
		if i == retryCount {
			log.Println("Connecting to Postgresql failed after multiple retries")
			return err
		}
		log.Println("Retrying connection to Postgresql in 10 seconds...")
		time.Sleep(10 * time.Second)
	}
	return nil
}

func ConnectToPostgresql() error {
	err := InitializeDBPool()
	if err != nil {
		log.Println("Connecting to PostgreSQL DB failed! ", err.Error())
		return err
	}
	_, err = DB().Exec(context.Background(), "DROP TABLE IF EXISTS messages;")
	if err != nil {
		log.Println("Failed to drop PostgreSQL table! ", err.Error())
		return err
	}

	_, err = DB().Exec(context.Background(), schema.WritingEntryTable)
	if err != nil {
		log.Println("Failed to create new PostgreSQL table! ", err.Error())
		return err
	}
	return nil
}
