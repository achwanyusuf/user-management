package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
	"net"
	"github.com/achwanyusuf/user-management/proto"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/achwanyusuf/user-management/repository"
	"github.com/achwanyusuf/user-management/model"
)

type server struct{}
var db *sql.DB
var rep repository.UserRepository

func init() {
	//Logging Formatter
	log.SetFormatter(&log.JSONFormatter{})
	file, _ := os.OpenFile("log/serverLogging.log",  os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetOutput(file)
	var logger = log.New()
	logger.Out = file
	log.SetLevel(log.DebugLevel)

	// Loading ENV variable
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
	}

	//Open Database Connection
	log.Info("Connecting to database")
    dbConnection := os.Getenv("DB_INLINE")
	db, err = sql.Open ("postgres", dbConnection)
    if err != nil {
        log.Fatal("Invalid DB config:", err)
	}
	err = db.Ping()
    if err != nil {
        log.Fatal("Database not connected")
    }
	log.Info("Database connected")
	rep = repository.NewRepositoryPostgres(db)
}

func main(){
	//Loading Env
	err := godotenv.Load()
	log.Info("Loading env")
    if err != nil {
        log.Error("Error loading .env file")
	}
	grpcPort := os.Getenv("GRPC_PORT")

	// Starting GRPC SERVER
	openPort, errOpenPort := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error(errOpenPort)
	}
	log.Info("listening port " + grpcPort)
	serverService := grpc.NewServer()
	log.Info("Starting GRPC server")
	proto.RegisterAddServiceServer(serverService, &server{})
	reflection.Register(serverService)
	serving := serverService.Serve(openPort)
	if serving != nil{
		log.Error(serving)
	}
	log.Info("GRPC Server is Started")
}


func userData (ur repository.UserRepository, input string) *model.UserData{
	userData, err := ur.ReadOneByEmail(input)
	if err != nil{
		return nil
	}
	return userData
}

func readAll (ur repository.UserRepository) *proto.ReadAllResponse{
	userDatas, _ := ur.ReadAll()
	return userDatas
}

func userDataById (ur repository.UserRepository, input string) *model.UserData{
	userData, err := ur.ReadOneByUserId(input)
	if err != nil{
		return nil
	}
	return userData
}

func readToken (ur repository.UserRepository, token string) bool{
	tokenValidity := ur.ReadToken(token)
	return tokenValidity
}

func updateUser (ur repository.UserRepository, data *model.UserData) bool{
	err := ur.Update(data)
	return err
}

func updateUser2 (ur repository.UserRepository, data *model.UserData) bool{
	err := ur.Update2(data)
	return err
}


func createUser (ur repository.UserRepository, data *model.UserData) bool{
	err := ur.Create(data)
	return err
}


func deleteUser (ur repository.UserRepository, userId string) bool{
	err := ur.Delete(userId)
	return err
}