package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/azmanabdlh/go-sample-api/internal/book"
	"github.com/azmanabdlh/go-sample-api/internal/httpx"
	"github.com/azmanabdlh/go-sample-api/internal/provider"
	"github.com/azmanabdlh/go-sample-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	r.POST("/echo", func(c *gin.Context) {
		// var body map[string]interface{}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			http.Error(c.Writer, "error", 500)
		}

		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.WriteHeader(200)
		c.Writer.Write(body)

	})

	repo := repository.NewMemoryStore()
	service := book.NewService(repo)
	handler := httpx.NewHandler(service)

	r.POST("/auth/token", handler.GenerateToken)

	r.Use(
		httpx.RequiredAuthentication(
			provider.NewMemoryTokenProvider(),
			// or json-web-token provider
			// provider.NewJsonWebTokenProvider(),
		),
	)
	{
		r.POST("/books", handler.Create)
		r.GET("/books", handler.Search)
		r.GET("/books/:id", handler.FindByID)
		r.PUT("/books/:id", handler.Update)
		r.DELETE("/books/:id", handler.Delete)

	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("http server running at :" + port)

		err := server.ListenAndServe()

		if err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// wait OS signal
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	sig := <-quit

	log.Printf("signal received: %v", sig)

	// shutdown timeout context
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	// graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Println("server stopped gracefully")

}
