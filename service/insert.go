package service

import (
	"context"
	"fmt"
	"log"
	"operate-mongodb/common"
)

func MongoInsert() {
	mongodb := common.CONFIG.MongoDb
	// 指定获取要操作的数据集
	collection := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collection)

	//使用collection.InsertMany()方法插入多条文档记录
	a := make([]interface{}, len(common.CONFIG.ClothTemplate), len(common.CONFIG.ClothTemplate))
	for i := range common.CONFIG.ClothTemplate {
		a[i] = common.CONFIG.ClothTemplate[i]
	}
	insertManyResult, err := collection.InsertMany(context.TODO(), a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

/*
//创建一些ClothTemplate类型的值，准备插入到数据库中
	s1 := model.ClothTemplate{
		Class:    "com.sweetpotato.model.template.ClothTemplate",
		Cid:      "1001",
		GetType:  1,
		Name:     "yaya经典款发型",
		Position: "[-4.504,124.629]",
		Scale:    "0.5",
		SuitId:   1,
		Type:     1,
		Price:    0,
	}
*/
