package config

type Database struct {
	Mysql *Mysql `mapstructure:"mysql"`
	Redis *Redis `mapstructure:"redis"`
}

type Mysql struct {
	Addr     string `mapstructure:"addr"`
	Port     int16  `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Port     int16  `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
