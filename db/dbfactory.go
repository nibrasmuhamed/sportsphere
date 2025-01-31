package db

import (
	"github.com/nibrasmuhamed/sportsphere/config"
	"github.com/nibrasmuhamed/sportsphere/pkg/constants"
	"github.com/nibrasmuhamed/sportsphere/pkg/helper"
)

func InitDatabase() {
	var db DatabaseClient
	url := ""
	if config.Get().DatabaseType == string(constants.MONGODB) {
		db = NewMongoDB()
		url = helper.GetEnv("MONGO_URL", "mongodb://localhost:27017")
	}

	err := db.Connect(url, config.Get().DatabaseName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	SetDatabase(db)
}
