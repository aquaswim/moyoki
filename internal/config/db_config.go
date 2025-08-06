package config

type DBConfig struct {
	DSN string `env:"DB_DSN,notEmpty"`
}
