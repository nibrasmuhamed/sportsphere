package db

import (
	"github.com/mohdjishin/sportsphere/config"
	"github.com/mohdjishin/sportsphere/pkg/helper"
)

func InitDatabase() {
	var db DatabaseClient
	if config.Get().DatabaseType == "mongodb" {
		db = NewMongoDB()
	}
	err := db.Connect(helper.GetEnv("MONGO_URI", "mongodb://localhost:27017"), config.Get().DatabaseName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	SetDatabase(db)
}
