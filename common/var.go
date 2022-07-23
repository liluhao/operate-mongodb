package common

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"operate-mongodb/model"
)

var (
	CONFIG model.Config
	VP     *viper.Viper
	Client *mongo.Client
)
