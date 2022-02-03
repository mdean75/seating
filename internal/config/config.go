// Package config defines the configuration details for the application. Use of interfaces provides for different types
// of configuration such as env variables or json file.
package config

// Loader is the interface implemented by types that provide a means to load config values.
type Loader interface {
	LoadConfig() Configuration
}

// Configuration holds the configuration model for the application
type Configuration struct {
	MongoConfig
	DebugCors bool
}

func NewConfiguration(m MongoConfig, debugCors bool) Configuration {
	return Configuration{m, debugCors}
}

type MongoConfig struct {
	dbConn string
	dbName string
	groupCollection string
	eventCollection string
}

func NewMongoConfig(conn, dbName, groupCol, eventCol string) MongoConfig {
	return MongoConfig{dbConn: conn, dbName: dbName, groupCollection: groupCol, eventCollection: eventCol}
}

func (m *MongoConfig) DBConn() string {
	return m.dbConn
}

func (m *MongoConfig) SetDBConn(dbConn string) {
	m.dbConn = dbConn
}