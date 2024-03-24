package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AfsanehHabibi/neveshtedan/graph"
	"github.com/AfsanehHabibi/neveshtedan/internal/auth"
	"github.com/AfsanehHabibi/neveshtedan/internal/module"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:5173"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	postgres.Setup()
	userRepo := postgres.NewUserRepository(postgres.DB())
	router.Use(auth.Middleware(userRepo))
	module := module.NewNeveshtedanModule(postgres.DB())
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ M: module}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
