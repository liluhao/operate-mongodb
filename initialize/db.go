package initialize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"operate-mongodb/common"
)

func InitMongoDb() {
	mongodb := common.CONFIG.MongoDb
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + mongodb.Path)

	if client, err := mongo.Connect(context.TODO(), clientOptions); err != nil { // 连接到MongoDB
		log.Fatal(err)
	} else if err = client.Ping(context.TODO(), nil); err != nil { // 检查连接
		log.Fatal(err)
	} else {
		common.Client = client
		fmt.Println("Connected to MongoDB!")
	}
}
