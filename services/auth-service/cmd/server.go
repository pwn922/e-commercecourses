package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/pwn922/auth-service/graph"
	"github.com/pwn922/auth-service/graph/middlewares"
	"github.com/pwn922/auth-service/internal/config"
	"github.com/pwn922/auth-service/internal/database"
	"github.com/pwn922/auth-service/internal/handlers"
	"github.com/pwn922/auth-service/internal/repositories"
	"github.com/pwn922/auth-service/internal/services/auth"
	"github.com/pwn922/auth-service/internal/services/jwt"
	"github.com/pwn922/auth-service/internal/services/user"
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
	userRepository := repositories.NewUserRepository(db)
	roleRepository := repositories.NewRoleRepository(db)
	userService := user.NewUserService(userRepository, roleRepository)
	jwtService := jwt.NewJWTService()
	authService := auth.NewAuthService(userService, jwtService)
	resolver := graph.NewResolver(authService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers:  resolver,
	}))

	router := mux.NewRouter()
	authHandlers := handlers.NewAuthHandlers(authService)
    router.HandleFunc("/login", authHandlers.LoginHandler)
	router.HandleFunc("/register", authHandlers.RegisterHandler)
	
	authMiddleware := middlewares.NewAuthMiddleware(jwtService)
	router.Handle("/query", authMiddleware.Authenticate(srv))
    router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}