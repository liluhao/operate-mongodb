package conf

import "operate-mongodb/model"

type Config struct {
	MongoDb    MongoDb            `json:"mongodb,omitempty" yaml:"mongodb,omitempty" toml:"mongodb,,omitempty" bson:"mongodb,omitempty"`
	Chart      []model.Chart      `json:"chart,omitempty" yaml:"chart,omitempty" toml:"chart,,omitempty" bson:"chart,omitempty"`
	DashBoard  []model.DashBoard  `json:"dashboard,omitempty" yaml:"dashboard,omitempty" toml:"dashboard,omitempty" bson:"dashboard,omitempty"`
	DataBase   []model.DataBase   `json:"database,omitempty" yaml:"database,omitempty" toml:"database,omitempty" bson:"database,omitempty"`
	DataSet    []model.DataSet    `json:"dataset,omitempty" yaml:"dataset,omitempty" toml:"dataset,omitempty" bson:"dataset,omitempty"`
	Transtable []model.Transtable `json:"transtable,omitempty" yaml:"transtable,omitempty" toml:"transtable,omitempty" bson:"transtable,omitempty"`
}

type MongoDb struct {
	//example: mongodb.Path=localhost:27017
	Path string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty" bson:"path,omitempty"`
	//example: dbname=game
	Dbname string `json:"dbname,omitempty" yaml:"dbname,omitempty" toml:"dbname,omitempty" bson:"dbname,omitempty"`
	//example: collection=clothTemplate
	Collections []string `json:"collection,omitempty" yaml:"collection,omitempty" toml:"collection,omitempty" bson:"collection,omitempty"`
}
