package config

import (
	"os"
	"strconv"
)

type EnvVar struct {
	dbConn            string
}

func (e EnvVar) LoadConfig() Configuration {
	e.dbConn = os.Getenv("dbconn")

	dc := os.Getenv("DEBUG_CORS")
	debugCors, err := strconv.ParseBool(dc)
	if err != nil {
		return NewConfiguration(NewMongoConfig(e.dbConn), false)
	}

	return NewConfiguration(NewMongoConfig(e.dbConn), debugCors)
}
