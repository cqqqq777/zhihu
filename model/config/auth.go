package config

type Auth struct {
	Jwt *Jwt `mapstructure:"jwt"`
}

type Jwt struct {
	ExpiresTime int    `mapstructure:"expires-time"`
	SecretKey   string `mapstructure:"secret-key"`
}
