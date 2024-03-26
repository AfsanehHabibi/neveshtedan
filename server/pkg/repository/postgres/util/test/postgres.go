package test

import (
	"context"
	"log"
	"os/exec"

	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func PreparePostgresForTest() *pgxpool.Pool {
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
	dropTables([]string{"writings", "writing_text_fields", "writing_number_fields", "writing_image_fields",
		"writing_video_fields", "users", "template_fields"})

	return postgres.DB()
}

func dropTables(names []string) {
	for _, name := range names {
		_, err := postgres.DB().Exec(context.Background(), "DROP TABLE IF EXISTS "+name+";")
		if err != nil {
			log.Fatalln("error in dropping table ", name, err.Error())
		}
	}
}
