package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"seating/internal/app/services/eventservice"
	"seating/internal/app/services/groupservice"
	"seating/internal/config"
	"seating/internal/db"
	eventadapter "seating/internal/handlers/event"
	groupadapter "seating/internal/handlers/group"
	"seating/internal/repositories/eventrepo"
	"seating/internal/repositories/grouprepo"
	"syscall"
	"time"
)

func NewHTTP(groupService *groupadapter.HTTPHandler, eventService *eventadapter.HTTPHandler) *http.Server{
	crs := NewRouterWithCors(groupService, eventService)


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
	conf.MongoConfig.SetDBConn("mongodb://127.0.0.1:27017") // need to remove
	
	// mongoClient, err := db.NewMongoDatabase()
	mongoConn, err := db.NewMongoDatabase(conf.DBConn())
	if err != nil {
		// return nil, err
		fmt.Println("unable to connect to mongo: ", err)
		return
	}

	groupRepo := grouprepo.NewDAO(mongoConn, "testdb", "group")
	eventrepo := eventrepo.NewDAO(mongoConn, "testdb", "event")

	groupService := groupservice.New(groupRepo)
	eventService := eventservice.New(eventrepo)

	groupHandler := groupadapter.NewHTTPHandler(groupService)
	eventHandler := eventadapter.NewHTTPHandler(eventService)

	// g, err := group.CreateController()
	// if err != nil {
	// 	fmt.Printf("unable to create controller, error: %v\n", err)
	// 	return
	// }

	// c := app.NewController(g)
	// if err != nil {
	// 	fmt.Printf("unable to create controller, error: %v\n", err)
	// 	return
	// }

	srv := NewHTTP(groupHandler, eventHandler)

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
		//log.Fatalf("Server Shutdown Failed:%+v", err)
		return
	}
	log.Print("Server Exited Properly")
}