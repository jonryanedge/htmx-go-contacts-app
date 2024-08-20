package main

import (
	"fmt"
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
	archive archiver.Archiver
	debug   bool
}

func main() {
	err := godotenv.Load()
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		isDebug = false
		fmt.Println("no debug")
	}
	addr := os.Getenv("ADDR")
	if "" == addr {
		addr = ":3333"
	}

	app := app{
		archive: *archiver.NewArchiver(),
		debug:   isDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// e.Logger.Fatal(e.Start(":3333"))
}
