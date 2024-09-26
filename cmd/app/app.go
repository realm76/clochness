package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/realm76/clochness/internal/app"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	host := "localhost"
	port := 5432
	user := "admin"
	password := "admin"
	dbName := "clochness"

	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	dbpool, err := pgxpool.New(context.Background(), connectString)
	if err != nil {
		panic(err)
	}

	defer dbpool.Close()

	if err := run(ctx, dbpool); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, db *pgxpool.Pool) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	defer cancel()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	appServer := app.NewServer(sugar, db)
	httpServer := appServer.HttpServer()

	go func() {
		log.Printf("server listening on %s", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)

		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
	}()

	wg.Wait()

	return nil
}
