type Config struct {
	Config     string `yaml:"env" env-default: "development`
	StoraePath string `taml:"storage_path" env-required: "true"`
	HTTPServer string `yaml:"http_server"`
}

type HTTPServer struct {
	Adress      string        `yaml:"adress" env-default:"0.0.0.:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default: "5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default: "60s"`
}