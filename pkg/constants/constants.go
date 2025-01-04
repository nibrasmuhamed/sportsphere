package constants

type EnvironmentVariable string

const (
	CONFIG_PATH EnvironmentVariable = "CONFIG_PATH"
	MONGO_URI   EnvironmentVariable = "MONGO_URI"
)

type DatabaseType string

const (
	MONGODB DatabaseType = "mongodb"
)
