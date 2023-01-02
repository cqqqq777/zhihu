package config

type App struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int16  `mapstructure:"port"`
	Version string `mapstructure:"version"`
}
