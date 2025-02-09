package main

import (
	"log"
	"os"
	"realtyV2/internal/data"
	"realtyV2/internal/scraper"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type Application struct {
	log     zerolog.Logger
	s       *echo.Echo
	scraper scraper.Scraper
	store   *data.Store
	// TODO: temporary
	cache map[string]BoundBox
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := newApp()
	// app.s.Use(middleware.Logger())
	app.routes()
	app.s.Start(":8383")
}

func newServer() *echo.Echo {
	s := echo.New()
	return s
}
func newApp() *Application {
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	}
	log := zerolog.New(os.Stdout).Level(zerolog.Level(logLevel)).With().Caller().Timestamp().Logger()
	store, err := data.NewStore(os.Getenv("POSTGRES_URL"), log)
	if err != nil {
		panic(err)
	}

	return &Application{
		log:     log,
		s:       newServer(),
		scraper: scraper.Scraper{Log: log, Size: 20},
		store:   store,
		cache:   make(map[string]BoundBox),
	}
}
