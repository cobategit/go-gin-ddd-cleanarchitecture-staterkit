package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PgConfig struct {
	DSN string
}

type MysqlConfig struct {
	DSN string
}

type JwtConfig struct {
	Secret string
	Issuer string
}

type Config struct {
	AppPort  string
	DBDriver string /* pgsql or mysql */

	Pgsql PgConfig
	Mysql MysqlConfig
	Jwt   JwtConfig
}

func getEnv(key, def string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func Load() *Config {
	cfg := &Config{
		AppPort:  getEnv("APP_PORT", "5005"),
		DBDriver: getEnv("DB_DRIVER", "postgres"),

		Pgsql: PgConfig{
			DSN: getEnv("POSTGRES_DSN", "host=postgres port=5432 user=postgres password= dbname=practice sslmode=disable"),
		},
		Mysql: MysqlConfig{
			DSN: getEnv("MYSQL_DSN", "root:password@tcp(mysql:3306)/app_db?parseTime=true"),
		},
		Jwt: JwtConfig{
			Secret: getEnv("JWT_SECRET", "supersecret"),
			Issuer: getEnv("JWT_ISSUER", "go-gin-ddd"),
		},
	}

	if cfg.Jwt.Secret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return cfg
}
