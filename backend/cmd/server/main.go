package main

import (
	"context"
	"cooking/backend/internal/config"
	"cooking/backend/internal/db/postgres"
	"cooking/backend/internal/handlers"
	recipeHandler "cooking/backend/internal/handlers/recipe"
	scheduleHandler "cooking/backend/internal/handlers/schedule"
	mm "cooking/backend/internal/middleware"
	postgresRepo "cooking/backend/internal/repository/postgres"
	"cooking/backend/internal/usecase"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := config.SetUp()
	if err != nil {
		slog.Error("failed to fetch config", "error", err)
	}
	postgresDB, err := postgres.InitDB()
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
	}
	err = postgres.MakeMigrations(true)
	if err != nil {
		slog.Error("failed to make migrations", "error", err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	userRepo := postgresRepo.NewUserRepo(postgresDB)
	//appointmentRepo := postgresRepo.NewAppointmentRepo(postgresDB)
	recipeRepo := postgresRepo.NewRecipeRepo(postgresDB)
	chefRepo := postgresRepo.NewChefRepo(postgresDB)
	scheduleRepo := postgresRepo.NewScheduleRepo(postgresDB)
	//subscriptionRepo := postgresRepo.NewSubscriptionRepo(postgresDB)

	userUse := usecase.UserInstance(userRepo)
	//appointmentUse := usecase.AppointmentInstance(appointmentRepo)
	recipeUse := usecase.RecipeInstance(recipeRepo)
	chefUse := usecase.ChefInstance(chefRepo)
	scheduleUse := usecase.ScheduleInstance(scheduleRepo)
	//subscriptionUse := usecase.SubscriptionInstance(subscriptionRepo)

	recipeHandlers := recipeHandler.NewRecipeHandler(recipeUse, userUse)
	scheduleHandlers := scheduleHandler.NewScheduleHandler(scheduleUse, chefUse)

	authHandler := handlers.NewAuthHandler(userUse, chefUse)
	e.POST("/api/register", authHandler.Register)
	e.POST("/api/login", authHandler.Authenticate)

	protectedGroup := e.Group("/api", mm.JwtMiddleware)

	recipeGroup := protectedGroup.Group("/recipes")
	recipeGroup.GET(":id/", recipeHandlers.GetRecipe)
	recipeGroup.POST("", recipeHandlers.CreateRecipe)
	recipeGroup.PUT(":id/", recipeHandlers.UpdateRecipe)
	recipeGroup.DELETE(":id/", recipeHandlers.DeleteRecipe)

	scheduleGroup := protectedGroup.Group("/schedules")
	scheduleGroup.GET(":id/", scheduleHandlers.GetEvent)
	scheduleGroup.POST("", scheduleHandlers.CreateEvent)
	scheduleGroup.PUT(":id/", scheduleHandlers.UpdateEvent)
	scheduleGroup.DELETE(":id/", scheduleHandlers.DeleteEvent)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":" + config.AppConfig.Server.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start server", "error", err)
		}
	}()

	<-stop
	slog.Info("received shutdown signal, starting shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Error("failed to gracefully shut down server", "error", err)
	}

	slog.Info("server gracefully stopped")
}
