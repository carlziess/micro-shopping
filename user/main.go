package main

import (
	"demo/micro/shopping/user/model"
	"demo/micro/shopping/user/repository"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"
	"log"
)

func main() {

	//加载配置项
	err := config.LoadFile("./config.json")
	if err != nil{
		log.Fatal("Could not load config file: %s",err.Error())
		return
	}
	conf := config.Map()

	//db
	db,err := CreateConnection(conf["mysql"].(map[string]interface{}))
	defer db.Close()

	db.AutoMigrate(&model.User{})

	if err != nil{
		log.Fatalf("connection error : %v \n",err)
	}

	repo := &repository.User{db}

	//New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	//Register Handler
	user.Regis







}
