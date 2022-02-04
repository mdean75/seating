package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"seating/internal/app/services/attendeeservice"
	"seating/internal/app/services/eventservice"
	"seating/internal/app/services/groupservice"
	"seating/internal/config"
	"seating/internal/db"
	"seating/internal/handlers/attendeeadapter"
	"seating/internal/handlers/eventadapter"
	"seating/internal/handlers/groupadapter"
	"seating/internal/repositories/attendeerepo"
	"seating/internal/repositories/eventrepo"
	"seating/internal/repositories/grouprepo"
	"syscall"
	"time"
)

func NewHTTP(groupService *groupadapter.HTTPHandler, eventService *eventadapter.HTTPHandler, attendeeService *attendeeadapter.HTTPHandler, conf config.Configuration) *http.Server{
	crs := NewRouterWithCors(groupService, eventService, attendeeService, conf)


	s := http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      crs,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &s
}

func Run() {
	stop := make(chan os.Signal, 1)
	kill := make(chan struct{}, 1)

	conf := config.EnvVar{}.LoadConfig()
	
	mongoConn, err := db.NewMongoDatabase(conf.DBConn())
	if err != nil {
		fmt.Println("unable to connect to mongo: ", err)
		return
	}

	groupRepo := grouprepo.NewDAO(mongoConn, "testdb", "group")
	eventrepo := eventrepo.NewDAO(mongoConn, "testdb", "event")
	attendeeRepo := attendeerepo.NewDAO(mongoConn, "testdb", "attendee")

	groupService := groupservice.New(groupRepo)
	eventService := eventservice.New(eventrepo)
	attendeeService := attendeeservice.New(attendeeRepo)

	groupHandler := groupadapter.NewHTTPHandler(groupService)
	eventHandler := eventadapter.NewHTTPHandler(eventService)
	attendeeHandler := attendeeadapter.NewHTTPHandler(attendeeService)

	srv := NewHTTP(groupHandler, eventHandler, attendeeHandler, conf)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	log.Print("Server Started: ", srv.Addr)

	select {
	case <-stop:
		fmt.Println("os stop signal received")
	case <-kill:
		fmt.Println("kill signal received")
	}

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		return
	}
	log.Print("Server Exited Properly")
}