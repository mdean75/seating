package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewHTTP() *http.Server{
	crs := NewRouterWithCors()


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

	srv := NewHTTP()

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	log.Print("Server Started")

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