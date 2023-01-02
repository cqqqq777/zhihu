package config

type Logger struct {
	MaxSize   int    `mapstructure:"max-size"`
	MaxGroups int    `mapstructure:"max-groups"`
	MaxAge    int    `mapstructure:"max-age"`
	SavePath  string `mapstructure:"save-path"`
	LogType   string `mapstructure:"log-type"`
}
