package env

import "os"

type Config struct {
	AppPort string
	DB      struct {
		Host, Port, User, Pass, Name, SSLMode, Timezone string
	}
}

func Load() Config {
	var c Config
	c.AppPort = getenv("APP_PORT", ":8080")

	c.DB.Host = getenv("DB_HOST", "localhost")
	c.DB.Port = getenv("DB_PORT", "5432")
	c.DB.User = getenv("DB_USER", "root")
	c.DB.Pass = getenv("DB_PASS", "root")
	c.DB.Name = getenv("DB_NAME", "root")
	c.DB.SSLMode = getenv("DB_SSLMODE", "disable")
	c.DB.Timezone = getenv("DB_TIMEZONE", "America/Sao_Paulo")

	return c
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
