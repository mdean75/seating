package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"seating/internal/app"
	"seating/internal/app/group"
	"syscall"
	"time"
)

func NewHTTP(controller *app.Controller) *http.Server{
	crs := NewRouterWithCors(controller)


	s := http.Server{
		Addr:         "0.0.0.0:4000",
		Handler:      crs,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &s
}

func Run() {
	stop := make(chan os.Signal, 1)
	kill := make(chan struct{}, 1)

	g, err := group.CreateController()
	if err != nil {
		fmt.Printf("unable to create controller, error: %v\n", err)
		return
	}

	c := app.NewController(g)
	if err != nil {
		fmt.Printf("unable to create controller, error: %v\n", err)
		return
	}

	srv := NewHTTP(c)

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