package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pwn922/users-service/graph"
	"github.com/pwn922/users-service/internal/config"
	"github.com/pwn922/users-service/internal/database"
	"github.com/pwn922/users-service/internal/repositories"
	"github.com/pwn922/users-service/internal/services"
)

const defaultPort = "8080"

func main() {
	cfg := config.LoadConfig() 
    database.InitDatabase(cfg)
    defer database.CloseDatabase()
    database.Migrate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.GetDB()
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)


	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserService: userService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
