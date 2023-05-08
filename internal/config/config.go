package config

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresDbHost   string
}

func NewConfig() Config {
	return Config{
		PostgresUser:     "test",
		PostgresPassword: "test",
		PostgresDb:       "test",
		PostgresDbHost:   "localhost",
	}
}
