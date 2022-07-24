package env

type Startup struct {
	Dir     string
	File    string
	Content string
}

type Config struct {
	Enabled bool  `json:"enabled"`
	Delay   int32 `json:"delay"`
}

type App struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Path    string `json:"path"`
}

type Configuration struct {
	Conf Config `json:"config"`
	Apps []App  `json:"apps"`
}
