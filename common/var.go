package common

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"operate-mongodb/conf"
)

var (
	CONFIG conf.Config
	VP     *viper.Viper
	Client *mongo.Client
)
