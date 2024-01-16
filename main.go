package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := &http.Server{
		Addr:    ":8181",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	sinalQueChegou := <-sc

	fmt.Println(" sinal que chegou", sinalQueChegou)

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		time.Sleep(time.Second * 10)
	// 		fmt.Println("n=", i)
	// 	}
	// }()

	// select {
	// case <-ctx.Done():
	// 	fmt.Println("bye")
	// case <-time.After(time.Second * 3):
	// 	fmt.Println("hello")
	// }

}
