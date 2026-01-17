package config

import "os"

type Config struct {
	Postgres Postgres
	GRPC     GRPC
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type GRPC struct {
	Addr string
}

func Load() Config {
	return Config{
		Postgres: Postgres{
			Host:     env("PG_HOST", "localhost"),
			Port:     env("PG_PORT", "5432"),
			User:     env("PG_USER", "postgres"),
			Password: env("PG_PASSWORD", ""),
			DBName:   env("PG_DB", "postgres"),
			SSLMode:  env("PG_SSLMODE", "disable"),
		},
		GRPC: GRPC{
			Addr: env("GRPC_ADDR", ":50051"),
		},
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
