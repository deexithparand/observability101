package config

import "os"

type ServerConfig struct {
	Host string
	Port string
}

func InitServerConfig() ServerConfig {

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	sc := ServerConfig{
		Host: host,
		Port: port,
	}

	return sc
}
