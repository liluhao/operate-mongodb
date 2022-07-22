package conf

import "operate-mongodb/model"

type Config struct {
	System     System             `mapstructure:"system,omitempty" bson:"system,omitempty"`
	MongoDb    MongoDb            `mapstructure:"mongodb,omitempty"bson:"mongodb,omitempty"`
	Chart      []model.Chart      `mapstructure:"chart,omitempty" bson:"chart,omitempty"`
	DashBoard  []model.DashBoard  `mapstructure:"dashboard,omitempty" bson:"dashboard,omitempty"`
	DataBase   []model.DataBase   `mapstructure:"database,omitempty" bson:"database,omitempty"`
	DataSet    []model.DataSet    `mapstructure:"dataset,omitempty" bson:"dataset,omitempty"`
	Transtable []model.Transtable `mapstructure:"transtable,omitempty" bson:"transtable,omitempty"`
}

type MongoDb struct {
	//example: mongodb.Path=localhost:27017
	Path string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty" bson:"path,omitempty"`
	//example: dbname=game
	Dbname string `json:"dbname,omitempty" yaml:"dbname,omitempty" toml:"dbname,omitempty" bson:"dbname,omitempty"`
	//example: collection=clothTemplate
	Collections []string `json:"collection,omitempty" yaml:"collection,omitempty" toml:"collection,omitempty" bson:"collection,omitempty"`
}
type System struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
