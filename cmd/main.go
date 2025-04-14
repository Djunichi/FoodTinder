// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:9000
// @BasePath /api/v1
package main

import (
	"context"
	"fmt"
	_ "food-tinder/docs"
	"food-tinder/internal/config"
	handler2 "food-tinder/internal/handler"
	"food-tinder/internal/migration"
	"food-tinder/internal/repository"
	"food-tinder/internal/service"
	"food-tinder/internal/worker"
	"food-tinder/migrations"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	conf, err := config.Load("config")
	if err != nil {
		log.Fatalf("Can not read config %v", err)
	}

	dbFiles := migrations.GetPostgresMigrations()
	dbVersion, err := migration.PostgresMigrate(conf.DB.URL, conf.Migration, dbFiles)
	if err != nil {
		log.Fatalf("Can not migrate db %v", err)
	}
	log.Printf("dbVersion: %v", dbVersion)

	db, err := repository.InitORM(conf.DB)
	if err != nil {
		log.Fatalf("Can not init db %v", err)
	}

	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := repository.NewMongoClient(mongoCtx, conf.MongoUrl)
	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	defer mongoClient.Disconnect(mongoCtx)

	repos := repository.NewRepositoryContainer(db, mongoClient)
	services := service.NewServiceContainer(repos)

	handler := handler2.NewHttpHandler(services, conf)
	handler.Init()

	f := worker.NewFeedFetcher(repos.Products)

	c := cron.New()
	_, err = c.AddFunc("0 0 * * *", f.FetchFeed(conf.FeedUrl))

	if err != nil {
		log.Fatalf("Can not update product feed %v", err)
	}

	c.Start()
	fmt.Println("Cron started. Waiting for daily task...")

	// setup graceful shutdown channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(sigChan)

	go func() {
		<-sigChan
		log.Println("Signal received. Initiating shutdown...")
		cancelFunc()
	}()

	<-ctx.Done()

	// Call Stop() with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	log.Println("Shutting down HTTP server...")
	if err := handler.Stop(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown error: %v", err)
	}

	log.Println("Stopping cron scheduler...")
	c.Stop()
}
