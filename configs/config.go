package configs

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecretKey string

	MidtransServerKey string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		JWTSecretKey:      os.Getenv("JWT_SECRET_KEY"),
		MidtransServerKey: os.Getenv("MIDTRANS_SERVER_KEY"),
	}

	return cfg
}
