package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"go.igmp.app/internal/archiver"

	"github.com/joho/godotenv"
)

type app struct {
	archive   archiver.Archiver
	debug     bool
	logger    *slog.Logger
	templates map[string]*template.Template
}

func main() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }

	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		isDebug = false
		fmt.Println("no debug")
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := app{
		archive:   *archiver.NewArchiver(),
		debug:     isDebug,
		logger:    logger,
		templates: templateCache,
	}

	srv := &http.Server{
		Addr:         addr+":"+port,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting server on IP:%s Port:%s", addr, port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
