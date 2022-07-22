package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"operate-mongodb/common"
	"operate-mongodb/initialize"
	"operate-mongodb/service"
)

func init() {
	initialize.InitConf()
	initialize.InitMongoDb()
}

func main() {
	defer close()
	router := gin.Default()
	router.GET("/hello", service.MongoInsert)
	if err := router.Run(":" + common.CONFIG.System.Port); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

func close() {
	//断开连接
	err := common.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
