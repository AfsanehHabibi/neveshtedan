package config

import (
	"fmt"
	"os"
	"strconv"
)

var PgConnStr = fmt.Sprintf(
	"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable ",
	os.Getenv("PG_DB_HOST"),
	os.Getenv("PG_DB_PORT"),
	os.Getenv("PG_DB_NAME"),
	os.Getenv("PG_DB_USER"),
	os.Getenv("PG_DB_PASSWORD"),
)

var defaultCPoolConf = map[string]int{
	"PG_MAX_CON":           50,
	"PG_MAX_CON_IDLE_TIME": 5,
	"CA_NUM_CON":           50,
	"CA_TIMEOUT":           5,
}

func GetConfigValue(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return defaultCPoolConf[key]
}
