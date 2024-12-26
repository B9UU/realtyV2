package main

import (
	"fmt"
	"log"
	"os"
	"realtyV2/internal/data"
	"realtyV2/internal/repo"
	"realtyV2/internal/scraper"
	"strconv"

	"github.com/jmoiron/sqlx"
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
	dd      *repo.Queries
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
	db, err := Newdb(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	return &Application{
		log:     log,
		s:       newServer(),
		scraper: scraper.Scraper{Log: log, Size: 20},
		store:   data.NewStore(db, log),
		dd:      repo.New(db),
		cache:   make(map[string]BoundBox),
	}
}

func Newdb(dbUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
