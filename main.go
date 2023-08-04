package main

// ref: https://www.golinuxcloud.com/hexagonal-architectural-golang/

import (
	"Messenger/internal/adapters/handler"
	"Messenger/internal/adapters/repository"
	"Messenger/internal/core/services"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	repo        = flag.String("db", "postgres", "Database for storing messages")
	redisHost   = "localhost:6379"
	httpHandler *handler.HTTPHandler
	svc         *services.MessengerService
)

func main() {
	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)
	switch *repo {
	case "redis":
		store := repository.NewMessengerRedisRepository(redisHost)
		svc = services.NewMessengerService(store)
	default:
		dbhost := os.Getenv("PGSQL_HOST")
		port := "5432"
		dbuser := os.Getenv("PGSQL_USER")
		dbpassword := os.Getenv("PGSQL_PW")
		dbname := os.Getenv("PGSQL_DB")
		store := repository.NewMessengerPostgresRepository(dbhost, port, dbuser, dbpassword, dbname)
		svc = services.NewMessengerService(store)
	}

	InitRoutes()

}

func InitRoutes() {
	router := gin.Default()
	httpHandler = handler.NewHTTPHandler(*svc)
	router.GET("/messages/:id", httpHandler.ReadMessage)
	router.GET("/messages", httpHandler.ReadMessages)
	router.POST("/messages", httpHandler.SaveMessage)
	router.Run(":3000")
}
