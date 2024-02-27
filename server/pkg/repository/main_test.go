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
)

var (
	ctx    = context.Background()
	wEImp  []repository.WritingEntryRepository
	wEFImp []repository.WritingEntryFieldRepository
)

func TestMain(m *testing.M) {
	wEFImp = append(wEFImp, inmemory.NewInMemoryWritingEntryFieldRepository())
	wEImp = append(wEImp, inmemory.NewInMemoryWritingEntryRepository(), preparePostgresForTest())
	m.Run()
}

func preparePostgresForTest() repository.WritingEntryRepository {
	cmd := exec.Command("docker", "run",  "-e", "POSTGRES_HOST_AUTH_METHOD=trust", 
	"-e", "POSTGRES_DB=broker", "-e", "POSTGRES_USER=admin",
	"-d","-p", "5432:5432", "postgres:latest")
    
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
	_, err = postgres.DB().Exec(context.Background(), schema.WritingEntryTable)
	if err != nil {
		log.Println("erffor", err.Error())
	}
	return postgres.NewPostgresWritingEntryRepository(postgres.DB())
}
