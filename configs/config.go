package configs

import (
	"github.com/caarlos0/env/v10"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

type conf struct {
	DBDriver      string `env:"DB_DRIVER"`
	DBHost        string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
	WebServerPort string `env:"WEB_SERVER_PORT"`
	JwtSecret     string `env:"JWT_SECRET"`
	JwtExpiresIn  int    `env:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() *conf {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg := conf{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)
	return &cfg
}
