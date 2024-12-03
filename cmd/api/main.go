package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type Application struct {
	log zerolog.Logger
	s   *echo.Echo
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := newApp()
	app.s.Use(middleware.Logger())
	app.routes()
	app.s.Start(":8383")
}

func newServer() *echo.Echo {
	s := echo.New()
	s.Validator = &CustomValidator{validator: validator.New()}
	return s
}
func newApp() *Application {
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	}
	return &Application{
		log: zerolog.New(os.Stdout).Level(zerolog.Level(logLevel)).With().Timestamp().Logger(),
		s:   newServer(),
	}
}
