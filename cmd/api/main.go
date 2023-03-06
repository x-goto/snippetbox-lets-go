package main

import (
	"database/sql"
	"flag"
	"goto/snippetbox/internal/data"
	"goto/snippetbox/internal/logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type config struct {
	addr string
	DB   struct {
		DSN         string
		MaxIdleCons int
		MaxOpenCons int
	}
}

type application struct {
	cfg    config
	logger *logger.Logger
	models data.Models
}

func main() {
	cfg := config{}
	logger := logger.New(os.Stdout, logger.LevelInfo)

	err := godotenv.Load()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("DB_DSN"), "connection string for database")
	flag.IntVar(&cfg.DB.MaxIdleCons, "db-max-idle-cons", 25, "sets max idle connections for the db")
	flag.IntVar(&cfg.DB.MaxOpenCons, "db-max-open-cons", 25, "sets max open connections for the db")
	flag.Parse()

	db, err := ConnectDB(cfg)
	if err != nil {
		logger.PrintFatal(err, map[string]string{
			"db": "error occured while connecting to database",
		})
	} else {
		logger.PrintInfo("database connection pool established", nil)
	}

	app := application{
		cfg:    cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	defer db.Close()
	err = app.serve()
	if err != nil {
		app.logger.PrintFatal(err, nil)
	}
}

func ConnectDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.DB.MaxIdleCons)
	db.SetMaxOpenConns(cfg.DB.MaxOpenCons)

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
