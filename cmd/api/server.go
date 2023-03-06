package main

import (
	"net/http"
	"os"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         app.cfg.addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.PrintInfo("", map[string]string{
		"addr": srv.Addr,
		"env":  os.Getenv("ENV"),
	})

	return srv.ListenAndServe()
}
