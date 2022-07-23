package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"operate-mongodb/common"
)

func MongoInsertOrUpdate(c *gin.Context) {
	var err error
	mongodb := common.CONFIG.MongoDb
	var ok = true
	opt := &options.ReplaceOptions{
		Upsert: &ok,
	}
	collection1 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[0])
	for i := range common.CONFIG.Chart {
		_, err = collection1.ReplaceOne(context.TODO(), bson.D{{"_id", common.CONFIG.Chart[i].Id}}, common.CONFIG.Chart[i], opt)
		if err != nil {
			log.Fatal(err)
		}
	}

	collection2 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[1])
	for i := range common.CONFIG.DashBoard {
		_, err = collection2.ReplaceOne(context.TODO(), bson.D{{"_id", common.CONFIG.Chart[i].Id}}, common.CONFIG.Chart[i], opt)
		if err != nil {
			log.Fatal(err)
		}
	}

	collection3 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[2])
	for i := range common.CONFIG.DataBase {
		_, err = collection3.ReplaceOne(context.TODO(), bson.D{{"_id", common.CONFIG.Chart[i].Id}}, common.CONFIG.Chart[i], opt)
		if err != nil {
			log.Fatal(err)
		}
	}

	collection4 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[3])
	for i := range common.CONFIG.DataSet {
		_, err = collection4.ReplaceOne(context.TODO(), bson.D{{"_id", common.CONFIG.Chart[i].Id}}, common.CONFIG.Chart[i], opt)
		if err != nil {
			log.Fatal(err)
		}
	}

	collection5 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[4])
	for i := range common.CONFIG.Transtable {
		_, err = collection5.ReplaceOne(context.TODO(), bson.D{{"_id", common.CONFIG.Chart[i].Id}}, common.CONFIG.Chart[i], opt)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("成功")
	c.JSON(http.StatusOK, gin.H{
		"message": "success！",
	})
}

/*
func MongoInsert(c *gin.Context) {
	mongodb := common.CONFIG.MongoDb
	// 指定获取要操作的数据集
	collection1 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[0])
	//使用collection.InsertMany()方法插入多条文档记录
	collect1 := make([]interface{}, len(common.CONFIG.Chart), len(common.CONFIG.Chart))
	for i := range common.CONFIG.Chart {
		//common.CONFIG.Chart[i].Id = primitive.NewObjectID().Hex()
		collect1[i] = common.CONFIG.Chart[i]
	}
	_, err := collection1.InsertMany(context.TODO(), collect1)
	if err != nil {
		log.Fatal(err)
	}

	collection2 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[1])
	collect2 := make([]interface{}, len(common.CONFIG.DashBoard), len(common.CONFIG.DashBoard))
	for i := range common.CONFIG.DashBoard {
		//common.CONFIG.DashBoard[i].Id = primitive.NewObjectID().Hex()
		collect2[i] = common.CONFIG.DashBoard[i]
	}
	_, err = collection2.InsertMany(context.TODO(), collect2)
	if err != nil {
		log.Fatal(err)
	}

	collection3 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[2])
	collect3 := make([]interface{}, len(common.CONFIG.DataBase), len(common.CONFIG.DataBase))
	for i := range common.CONFIG.DataBase {
		//common.CONFIG.DataBase[i].Id = primitive.NewObjectID().Hex()
		collect3[i] = common.CONFIG.DataBase[i]
	}
	_, err = collection3.InsertMany(context.TODO(), collect3)
	if err != nil {
		log.Fatal(err)
	}

	collection4 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[3])
	collect4 := make([]interface{}, len(common.CONFIG.DataSet), len(common.CONFIG.DataSet))
	for i := range common.CONFIG.DataSet {
		//common.CONFIG.DataSet[i].Id = primitive.NewObjectID().Hex()
		collect4[i] = common.CONFIG.DataSet[i]
	}
	_, err = collection4.InsertMany(context.TODO(), collect4)
	if err != nil {
		log.Fatal(err)
	}

	collection5 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[4])
	collect5 := make([]interface{}, len(common.CONFIG.Transtable), len(common.CONFIG.Transtable))
	for i := range common.CONFIG.Transtable {
		//common.CONFIG.Transtable[i].Id = primitive.NewObjectID().Hex()
		collect5[i] = common.CONFIG.Transtable[i]
	}
	_, err = collection5.InsertMany(context.TODO(), collect5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("插入成功")
	c.JSON(http.StatusOK, gin.H{
		"message": "insert success！",
	})
}
*/
