package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Port     int    `mapstructure:"DB_PORT"`
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Name     string `mapstructure:"DB_NAME"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type ServerConfig struct {
	Port int `mapstructure:"API_PORT"`
}

var AppConfig Config
