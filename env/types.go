package env

type Startup struct {
	Dir     string
	File    string
	Content string
}

type Config struct {
	Enabled bool  `mapstructure:"enabled"`
	Delay   int32 `mapstructure:"delay"`
}

type App struct {
	Enabled bool   `mapstructure:"enabled"`
	Name    string `mapstructure:"name"`
	Path    string `mapstructure:"path"`
}

type Configuration struct {
	Conf Config `mapstructure:"config"`
	Apps []App  `mapstructure:"apps"`
}