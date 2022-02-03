package config

import (
	"os"
	"strconv"
)

type EnvVar struct {
	dbConn            string
	dbName string
	groupCollection string
	eventCollection string
}

func (e EnvVar) LoadConfig() Configuration {
	e.dbConn = os.Getenv("DBCONN")
	e.dbName = os.Getenv("DB_NAME")
	e.groupCollection = os.Getenv("GROUP_COLLECTION")
	e.eventCollection = os.Getenv("EVENT_COLLECTION")

	dc := os.Getenv("DEBUG_CORS")
	debugCors, err := strconv.ParseBool(dc)
	if err != nil {
		return NewConfiguration(NewMongoConfig(e.dbConn, e.dbName, e.groupCollection, e.eventCollection), false)
	}

	return NewConfiguration(NewMongoConfig(e.dbConn, e.dbName, e.groupCollection, e.eventCollection), debugCors)
}
