package config

// Config structure
type Config struct {
	HTTP     HTTPConfiguration
	Database DatabaseConfiguration
}

// HTTPConfiguration - HTTP type configuration of env
type HTTPConfiguration struct {
	Host string `env:"API_HOST"`
	Port string `env:"API_PORT"`
}

// DatabaseConfiguration - Database type configuration of env
type DatabaseConfiguration struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Auth     string `env:"DATABASE_AUTH"`
}
