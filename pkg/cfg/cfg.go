package cfg

type Database struct {
	DSN string `env:"DSN" env-default:""`
}

type API struct {
	Address string `env:"ADDRESS" env-default:"0.0.0.0:8080"`
	Mode    string `env:"MODE" env-default:"debug"`
}
