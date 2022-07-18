package conf

import "operate-mongodb/model"

type Config struct {
	MongoDb       MongoDb               `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	ClothTemplate []model.ClothTemplate `mapstructure:"clothtemplate" json:"clothtemplate" yaml:"clothtemplate"`
}
type MongoDb struct {
	//example: mongodb.Path=localhost:27017
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	//example: dbname=game
	Dbname string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	//example: collection=clothTemplate
	Collection string `mapstructure:"collection" json:"collection" yaml:"collection"`
}
