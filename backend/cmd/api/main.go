package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/services"
	"log"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	cfg := config.LoadConfig()
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)

	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to Migrate database:", err)
	}
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)
	r := routes.SetupRoutes(authHandler)

	log.Println("Server running on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
