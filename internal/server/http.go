package server

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"

	//"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"seating/internal/app"
	"seating/internal/app/services/attendeeservice"
	"seating/internal/app/services/eventservice"
	"seating/internal/app/services/groupservice"
	"seating/internal/app/services/industryservice"
	"seating/internal/config"
	"seating/internal/db"
	"seating/internal/handlers/attendeeadapter"
	"seating/internal/handlers/eventadapter"
	"seating/internal/handlers/groupadapter"
	"seating/internal/handlers/industryadapter"
	"seating/internal/repositories/attendeerepo"
	"seating/internal/repositories/eventrepo"
	"seating/internal/repositories/grouprepo"
	"seating/internal/repositories/industryrepo"
	"syscall"
	"time"
)

func NewHTTP(controller *app.Controller, conf config.Configuration) *http.Server {
	crs := NewRouterWithCors(controller, conf)

	s := http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      crs,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &s
}

func Run() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

	stop := make(chan os.Signal, 1)
	kill := make(chan struct{}, 1)

	conf := config.EnvVar{}.LoadConfig()

	log.Debug().Msg("Connecting to database")
	mongoConn, err := db.NewMongoDatabase(conf.DBConn())
	if err != nil {
		fmt.Println("unable to connect to mongo: ", err)
		return
	}

	groupRepo := grouprepo.NewDAO(mongoConn, "testdb", "group")
	eventRepo := eventrepo.NewDAO(mongoConn, "testdb", "event")
	attendeeRepo := attendeerepo.NewDAO(mongoConn, "testdb", "event")
	industryRepo := industryrepo.NewDAO(mongoConn, "testdb", "industry")

	groupService := groupservice.New(groupRepo)
	eventService := eventservice.New(eventRepo)
	attendeeService := attendeeservice.New(attendeeRepo)
	industryService := industryservice.New(industryRepo)

	groupHandler := groupadapter.NewHTTPHandler(groupService)
	eventHandler := eventadapter.NewHTTPHandler(eventService)
	attendeeHandler := attendeeadapter.NewHTTPHandler(attendeeService)
	industryHandler := industryadapter.NewHTTPHandler(industryService)

	appController := app.NewController(groupHandler, eventHandler, attendeeHandler, industryHandler)

	log.Info().Msg("Start HTTP server")
	srv := NewHTTP(appController, conf)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	log.Info().Msgf("Server Started: ", srv.Addr)

	select {
	case <-stop:
		fmt.Println("os stop signal received")
	case <-kill:
		fmt.Println("kill signal received")
	}

	log.Info().Msg("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		return
	}
	log.Info().Msg("Server Exited Properly")
}
