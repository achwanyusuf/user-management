package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/achwanyusuf/user-management/proto"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, _ := os.OpenFile("log/clientLogging.log",  os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetOutput(file)
	var logger = log.New()
	logger.Out = file
	log.SetLevel(log.DebugLevel)
}

func main() {
	err := godotenv.Load()
	log.Info("Loading env")
    if err != nil {
        log.Error("Error loading .env file")
	}
	grpcPort := os.Getenv("GRPC_PORT")
	log.Info(grpcPort)
	clientAddress := os.Getenv("CLIENT_ADDRESS")
	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Starting Client Service")
	client := proto.NewAddServiceClient(conn)
	ginService := gin.Default()
	ginService.POST("/login", loginProcess(client))
	ginService.POST("/create", createProcess(client))
	ginService.POST("/delete", deleteProcess(client))
	ginService.POST("/update", updateProcess(client))
	ginService.POST("/readOne", readOneProcess(client))
	ginService.POST("/readAll", readAllProcess(client))
	startingGin := ginService.Run(clientAddress)
	if startingGin != nil{
		log.Fatal("Failed to start client: %v", startingGin)
	}
	log.Info("Client service is started")
}
