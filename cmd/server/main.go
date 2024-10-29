package main

import (
	"context"
	"os"

	"github.com/RhoNit/doctor_appointment_system/internal/db"
	handler "github.com/RhoNit/doctor_appointment_system/internal/handlers"
	"github.com/RhoNit/doctor_appointment_system/internal/repositories"
	"github.com/RhoNit/doctor_appointment_system/internal/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal("Error while initializing zap logger")
	}
	defer logger.Sync()

	dbConn, err := db.InitPgxConn(logger)
	if err != nil {
		logger.Fatal("Error while initializing the DB connection")
	}
	defer dbConn.Close(context.Background())

	docRepo := repositories.NewDoctorRepository(dbConn)
	docService := services.NewDoctorService(docRepo)
	docHandler := handler.NewDoctorHandler(docService, logger)

	e := echo.New()

	e.POST("/doctors", docHandler.CreateDoctorProfile)
	e.GET("/doctors/:doctor_id", docHandler.GetDoctorProfile)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		e.Start(":" + "8080")
	} else {
		e.Start(":" + serverPort)
	}
}
