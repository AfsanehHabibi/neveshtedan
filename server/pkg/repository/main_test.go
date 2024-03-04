package repository_test

import (
	"context"
	"log"
	"testing"

	"os/exec"

	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	inmemory "github.com/AfsanehHabibi/neveshtedan/pkg/repository/in_memory"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx     = context.Background()
	wEImp   []repository.WritingEntryRepository
	wEFImp  []repository.WritingEntryFieldRepository
	userRep repository.UserRepository
)

func TestMain(m *testing.M) {
	con := preparePostgresForTest()
	pWE := postgres.NewPostgresWritingEntryRepository(con)
	pWEF := postgres.NewPostgresWritingEntryFieldRepository(con)
	userRep = postgres.NewUserRepository(con)
	wEFImp = append(wEFImp, inmemory.NewInMemoryWritingEntryFieldRepository(), pWEF)
	wEImp = append(wEImp, inmemory.NewInMemoryWritingEntryRepository(), pWE)
	m.Run()
}

func preparePostgresForTest() *pgxpool.Pool {
	cmd := exec.Command("docker", "run", "-e", "POSTGRES_HOST_AUTH_METHOD=trust",
		"-e", "POSTGRES_DB=broker", "-e", "POSTGRES_USER=admin",
		"-d", "-p", "5432:5432", "postgres:latest")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error running Docker command:", err)
	}

	containerID := string(output)
	log.Println("Container ID:", containerID)
	err = postgres.InitializeDBPool()
	if err != nil {
		log.Fatalln("Connecting to DB failed! ", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), "DROP TABLE IF EXISTS writings;")
	if err != nil {
		log.Println("eror", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), "DROP TABLE IF EXISTS writing_fields;")
	if err != nil {
		log.Println("eror", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), "DROP TABLE IF EXISTS users;")
	if err != nil {
		log.Println("eror", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), schema.WritingEntryTable)
	if err != nil {
		log.Println("erffor", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), schema.WritingEntryFieldTable)
	if err != nil {
		log.Println("erffor", err.Error())
	}
	_, err = postgres.DB().Exec(context.Background(), schema.UserTable)
	if err != nil {
		log.Println("erffor", err.Error())
	}
	return postgres.DB()
}
