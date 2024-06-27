package main

import (
	"context"
	"example.com/demo"
	"example.com/demo/echo"
	"example.com/demo/mongo"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config, err := demo.NewConfig(ctx)

	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db := mongo.NewDB(config.Mongo, log.New(os.Stdout, "mongo: ", log.LstdFlags))

	err = db.Open()

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	userRepo := mongo.NewUserRepository(db)

	// Create a new echo server.
	server := echo.NewServer(echo.Options{
		EchoConfig: config.Echo,
		UserRepo:   userRepo,
	})

	server.RegisterHandlers()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	// Run the server.
	server.Run(ctx, wg)

	wg.Wait()

	log.Println("Gracefully exited server...")
}
