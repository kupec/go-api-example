package api

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	ApiBind     string `env:"API_BIND" envDefault:":8000"`
}
