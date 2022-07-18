package main

import (
	"context"
	"fmt"
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
	service.MongoInsert()
}

func close() {
	// 断开连接
	err := common.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
