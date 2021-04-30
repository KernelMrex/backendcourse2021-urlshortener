package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlshortener/pkg/urlshortener/app/service"
	"urlshortener/pkg/urlshortener/infrastructure/sqlite/migration"
	"urlshortener/pkg/urlshortener/infrastructure/sqlite/query"
	"urlshortener/pkg/urlshortener/infrastructure/sqlite/repository"
	"urlshortener/pkg/urlshortener/infrastructure/transport"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	conn, err := sqlx.Connect("sqlite3", os.Getenv("DATA_SOURCE_URL"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not connect to a database"))
		return
	}

	if err := migration.MigrateDB(conn, os.Getenv("MIGRATION_FILE_PATH")); err != nil {
		log.Fatal(errors.Wrap(err, "could not migrate a database"))
		return
	}

	redirectRepo := repository.NewRedirectRepository(conn)
	redirectService := service.NewRedirectService(redirectRepo)

	redirectQueryService := query.NewRedirectQueryService(conn)

	server := startServer(os.Getenv("SERVE_REST_ADDRESS"), transport.NewServer(redirectService, redirectQueryService))
	waitForKillSignal(getKillSignalChan())
	shutdownServer(server)
}

func startServer(serveURL string, srv *transport.Server) *http.Server {
	router := transport.NewRouter(srv)
	server := http.Server{
		Addr:    serveURL,
		Handler: router,
	}

	go func() {
		log.Info("Server is starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(errors.Wrap(err, "error while serving HTTP"))
		}
	}()

	return &server
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

func shutdownServer(server *http.Server) {
	log.Info("Server is shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
