package main

import (
	"log"
	"os"
	"strings"

	"pos-backend/internal/api"
	"pos-backend/internal/database"
	"pos-backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// SQLite database path — Tauri sets ALTEQIA_DB_PATH to the app data directory
	dbPath := getEnv("ALTEQIA_DB_PATH", "./pos.db")

	db, err := database.Connect(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := database.Ping(db); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	gin.SetMode(getEnv("GIN_MODE", "release"))
	router := gin.New()
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	defaultOrigins := []string{
		"http://localhost:3000", "http://localhost:3001", "http://localhost:3002",
		"http://localhost:3003", "http://localhost:5173",
		"http://192.168.1.59:5173",
		"tauri://localhost",          // Tauri v2 webview (Windows/Linux)
		"https://tauri.localhost",    // Tauri v2 webview (macOS)
	}
	extraOrigins := getEnv("CORS_ORIGINS", "")
	if extraOrigins != "" {
		for _, o := range strings.Split(extraOrigins, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				defaultOrigins = append(defaultOrigins, o)
			}
		}
	}
	allowedOrigins := defaultOrigins

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if extraOrigins == "*" {
				return true
			}
			for _, allowed := range allowedOrigins {
				if allowed == origin {
					return true
				}
			}
			return false
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	authMiddleware := middleware.AuthMiddleware()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "message": "POS API is running"})
	})

	apiRoutes := router.Group("/api/v1")
	api.SetupRoutes(apiRoutes, db, authMiddleware)

	// Default port 17432 to avoid conflicts with common services
	port := getEnv("PORT", "17432")
	log.Printf("Starting server on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
