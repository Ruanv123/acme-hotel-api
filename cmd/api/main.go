package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/ruanv123/acme-hotel-api/internal/api/handlers"
	"github.com/ruanv123/acme-hotel-api/internal/database"
	"github.com/ruanv123/acme-hotel-api/internal/logger"
	"github.com/ruanv123/acme-hotel-api/internal/middleware"
	"github.com/ruanv123/acme-hotel-api/internal/repository"
	"github.com/ruanv123/acme-hotel-api/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: error loading .env file: %s\n", err)
	}

	// inicializando a conexão com o banco
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying *sql.DB instance:", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	userRepo := repository.NewUserRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	authService := service.NewAuthService(
		userRepo,
		jwtSecret,
	)

	authHandler := handlers.NewAuthHandler(authService)

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	// public routes
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// API routes (protected)
	// apiRouter := router.PathPrefix("/api/v1").Subrouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-API-Key",
			"*", // Allow all headers
		},
		ExposedHeaders: []string{
			"Link",
		},
		AllowCredentials: false, // Must be false when using AllowedOrigins: ["*"]
		MaxAge:           300,
	})

	srv := &http.Server{
		Handler:      corsMiddleware.Handler(router),
		Addr:         ":" + getPort(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server
	logger.LogEvent(logrus.InfoLevel, "API started", logrus.Fields{
		"port": "8080",
	})
	log.Fatal(srv.ListenAndServe())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	return port
}
