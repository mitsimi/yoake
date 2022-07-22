package env

type Startup struct {
	Dir     string
	File    string
	Content string
}

type App struct {
	Enabled bool
	Name    string
	Path    string
}

type Config struct {
	Enabled bool
	Delay   int32
}
