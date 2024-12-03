package main

import (
	"log"
	"os"
	"realtyV2/internal/scraper"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type Application struct {
	log     zerolog.Logger
	s       *echo.Echo
	scraper scraper.Scraper
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
	log := zerolog.New(os.Stdout).Level(zerolog.Level(logLevel)).With().Timestamp().Logger()

	return &Application{
		log:     log,
		s:       newServer(),
		scraper: scraper.Scraper{Log: log},
	}
}
