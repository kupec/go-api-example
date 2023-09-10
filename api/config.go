package api

type Config struct {
	DatabaseDriver string `env:"DATABASE_DRIVER,required"`
	DatabaseSource string `env:"DATABASE_SOURCE,required"`
	ApiBind        string `env:"API_BIND" envDefault:":8000"`
}
