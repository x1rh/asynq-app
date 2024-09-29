package config

type Config struct {
	RedisConfig RedisConfig
}

type RedisConfig struct {
	Addr      string
	Password  string
	Username  string
	DB        int
	EnableTls bool
}
