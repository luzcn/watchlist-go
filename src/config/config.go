package config


type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
	AppEnv      string `env:"APP_ENV,default=development"`
}