package config

type Config struct {
	Database *Database
	Logger   *Logger
	Auth     *Auth
	App      *App
}
