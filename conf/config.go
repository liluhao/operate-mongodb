package conf

import "operate-mongodb/model"

type Config struct {
	MongoDb    MongoDb            `json:"mongodb" yaml:"mongodb" toml:"mongodb" bson:"mongodb"`
	Chart      []model.Chart      `json:"chart" yaml:"chart" toml:"chart" bson:"chart"`
	DashBoard  []model.DashBoard  `json:"dashboard" yaml:"dashboard" toml:"dashboard" bson:"dashboard"`
	DataBase   []model.DataBase   `json:"database" yaml:"database" toml:"database" bson:"database"`
	DataSet    []model.DataSet    `json:"dataset" yaml:"dataset" toml:"dataset" bson:"dataset"`
	Transtable []model.Transtable `json:"transtable" yaml:"transtable" toml:"transtable" bson:"transtable"`
	//ClothTemplate []model.ClothTemplate `json:"clothtemplate" yaml:"clothtemplate" toml:"clothtemplate" bson:"clothtemplate"`
}
type MongoDb struct {
	//example: mongodb.Path=localhost:27017
	Path string `json:"path" yaml:"path" toml:"path" bson:"path"`
	//example: dbname=game
	Dbname string `json:"dbname" yaml:"dbname" toml:"dbname" bson:"dbname"`
	//example: collection=clothTemplate
	Collections []string `json:"collection" yaml:"collection" toml:"collection" bson:"collection"`
}
