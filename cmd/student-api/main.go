package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aminshahid573/student-api/internal/config"
	"github.com/aminshahid573/student-api/internal/http/handlers/student"
	"github.com/aminshahid573/student-api/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started",
		slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Faild to Start Server")
		}
	}()

	<-done

	slog.Info("Shutting Down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to ShutDown Server", slog.String("error", err.Error()))
	}

	//you can also do this

	// if err != nil {
	// 	slog.Error("Failed to ShutDown Server", slog.String("error", err.Error()))
	// }

	slog.Info("Server ShutDown Successfully")

}
