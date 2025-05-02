package cfg

type Database struct {
	DSN string `env:"DSN" env-default:"host=localhost user=postgres password=postgres dbname=postgres port=5432"`
}

type API struct {
	Address string `env:"ADDRESS" env-default:"localhost:8080"`
	Mode    string `env:"MODE" env-default:"debug"`
}

type Auth struct {
	URL           string `env:"URL" env-default:"http://localhost:8081"`
	Realm         string `env:"REALM" env-default:"master"`
	AdminUsername string `env:"ADMIN_USERNAME" env-default:"admin"`
	AdminPassword string `env:"ADMIN_PASSWORD" env-default:"admin"`
	ClientID      string `env:"CLIENT_ID" env-default:"locpack-backend"`
	ClientSecret  string `env:"CLIENT_SECRET" env-default:"puKuYwCPLb7jUmbzrKMJCCCZcMLfU4Oy"`
}
