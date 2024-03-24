package repository_test

import (
	"context"
	"log"
	"testing"

	"os/exec"

	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx               = context.Background()
	writingEntry      repository.WritingEntryRepository
	writingEntryField repository.WritingEntryFieldRepository
	userRep           repository.UserRepository
	templateRep       repository.WritingTemplateRepository
	templateFieldRep  repository.WritingTemplateFieldRepository
)

func TestMain(m *testing.M) {
	con := preparePostgresForTest()
	writingEntry = postgres.NewPostgresWritingEntryRepository(con)
	writingEntryField = postgres.NewPostgresWritingEntryFieldRepository(con)
	userRep = postgres.NewUserRepository(con)
	templateRep = postgres.NewPostgresWritingTemplateRepository(con)
	templateFieldRep = postgres.NewPostgresWritingTemplateFieldRepository(con)
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
	_, err = postgres.DB().Exec(context.Background(), "DROP TABLE IF EXISTS template_fields;")
	if err != nil {
		log.Println("eror", err.Error())
	}
	return postgres.DB()
}
