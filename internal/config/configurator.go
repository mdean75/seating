package config

type DBConfigurator interface {
	DBConn() string
	SetDBConn(s string)
}
