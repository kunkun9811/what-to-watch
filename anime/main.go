package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	r := createRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error: HTTP Server shutdown error")
		}
	}()

	fmt.Println("Services instantiated. Start listening...")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	sig := <-c

	fmt.Printf("signal = %+v : Got interrupt signal, aborting...", sig)

	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)

	if err := server.Shutdown(timeoutCtx); err != nil {
		fmt.Println("Timeout when shutting down the server")
	}

	cancel()
	os.Exit(0)
}
