package main

import (
	"fileProcessing/config"
	"fileProcessing/internal/core/services"
	"fileProcessing/internal/handler"
	"fileProcessing/internal/repositories/raft"
	"fileProcessing/internal/repositories/redis"
	"fileProcessing/pkg/logging"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadConfig()

	logger := logging.NewLogger()
	defer logger.Close()

	conf := config.AppConfig{}

	redisRepo, _ := redis.NewRedisRepository()
	raftRepo := raft.NewRaftCluster(conf, redisRepo)
	raftNode := raftRepo.CreateNewRaftCluster()

	fileService := services.NewFileService(redisRepo)

	router := handler.NewRoutes(raftNode, fileService)

	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
