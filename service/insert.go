package service

import (
	"context"
	"log"
	"operate-mongodb/common"
)

func MongoInsert() {
	mongodb := common.CONFIG.MongoDb
	// 指定获取要操作的数据集
	collection1 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[0])
	//使用collection.InsertMany()方法插入多条文档记录
	collect1 := make([]interface{}, len(common.CONFIG.Chart), len(common.CONFIG.Chart))
	for i := range common.CONFIG.Chart {
		collect1[i] = common.CONFIG.Chart[i]
	}
	_, err := collection1.InsertMany(context.TODO(), collect1)
	if err != nil {
		log.Fatal(err)
	}

	collection2 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[1])
	collect2 := make([]interface{}, len(common.CONFIG.DashBoard), len(common.CONFIG.DashBoard))
	for i := range common.CONFIG.DashBoard {
		collect2[i] = common.CONFIG.DashBoard[i]
	}
	_, err = collection2.InsertMany(context.TODO(), collect2)
	if err != nil {
		log.Fatal(err)
	}

	collection3 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[2])
	collect3 := make([]interface{}, len(common.CONFIG.DataBase), len(common.CONFIG.DataBase))
	for i := range common.CONFIG.DataBase {
		collect3[i] = common.CONFIG.DataBase[i]
	}
	_, err = collection3.InsertMany(context.TODO(), collect3)
	if err != nil {
		log.Fatal(err)
	}

	collection4 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[3])
	collect4 := make([]interface{}, len(common.CONFIG.DataSet), len(common.CONFIG.DataSet))
	for i := range common.CONFIG.DataSet {
		collect4[i] = common.CONFIG.DataSet[i]
	}
	_, err = collection4.InsertMany(context.TODO(), collect4)
	if err != nil {
		log.Fatal(err)
	}

	collection5 := common.Client.Database(mongodb.Dbname).Collection(mongodb.Collections[4])
	collect5 := make([]interface{}, len(common.CONFIG.Transtable), len(common.CONFIG.Transtable))
	for i := range common.CONFIG.Transtable {
		collect5[i] = common.CONFIG.Transtable[i]
	}
	_, err = collection5.InsertMany(context.TODO(), collect5)
	if err != nil {
		log.Fatal(err)
	}

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
