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
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger *zap.SugaredLogger

func main() {
	logg, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize zap logger: %v", err))
	}
	defer logg.Sync()
	logger = logg.Sugar()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	conf := mustLoadConfig(getConfigPath())
	mustMigratePostgres(conf)
	db := mustInitPostgres(conf)
	mongoClient := mustInitMongo(conf)
	defer mongoClient.Disconnect(context.Background())

	repos := repository.NewRepositoryContainer(db, mongoClient, logger)
	services := service.NewServiceContainer(repos)
	handler := handler2.NewHttpHandler(services, conf, logger)
	handler.Init()

	scheduler := startCronJob(repos, conf)
	defer scheduler.Stop()

	setupGracefulShutdown(ctx, cancelFunc, handler)
}

func mustLoadConfig(path string) *config.Config {
	conf, err := config.Load(path)
	if err != nil {
		logger.Fatalf("Cannot read config: %v", err)
	}

	mongoUrl := os.Getenv("MONGO_URI")
	dbUrl := os.Getenv("DATABASE_URL")

	if mongoUrl != "" {
		conf.MongoUrl = mongoUrl
	}
	if dbUrl != "" {
		conf.DB.URL = dbUrl
	}

	return conf
}

func getConfigPath() string {
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		return path
	}
	return "config"
}

func mustMigratePostgres(conf *config.Config) {
	dbFiles := migrations.GetPostgresMigrations()
	dbVersion, err := migration.PostgresMigrate(conf.DB.URL, conf.Migration, dbFiles)
	if err != nil {
		logger.Fatalf("Cannot migrate db: %v", err)
	}
	logger.Infof("dbVersion: %v", dbVersion)
}

func mustInitPostgres(conf *config.Config) *gorm.DB {
	db, err := repository.InitORM(conf.DB)
	if err != nil {
		logger.Fatalf("Cannot init db: %v", err)
	}
	return db
}

func mustInitMongo(conf *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := repository.NewMongoClient(ctx, conf.MongoUrl)
	if err != nil {
		logger.Fatalf("Mongo connection failed: %v", err)
	}
	return client
}

func startCronJob(repos *repository.Container, conf *config.Config) *cron.Cron {
	f := worker.NewFeedFetcher(repos.Products)
	c := cron.New()
	_, err := c.AddFunc(fmt.Sprintf("%s %s %s %s %s", conf.Worker.Minute, conf.Worker.Hour, conf.Worker.Day, conf.Worker.Month, conf.Worker.DayOfWeek), f.FetchFeed(conf.FeedUrl))
	if err != nil {
		logger.Fatalf("Cannot schedule product feed update: %v", err)
	}
	c.Start()
	logger.Info("Cron started. Waiting for daily task...")
	return c
}

func setupGracefulShutdown(ctx context.Context, cancelFunc context.CancelFunc, handler handler2.HttpHandler) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(sigChan)

	go func() {
		<-sigChan
		logger.Info("Signal received. Initiating shutdown...")
		cancelFunc()
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	logger.Info("Shutting down HTTP server...")
	if err := handler.Stop(shutdownCtx); err != nil {
		logger.Fatalf("HTTP server shutdown error: %v", err)
	}

	logger.Info("Shutdown complete.")
}
