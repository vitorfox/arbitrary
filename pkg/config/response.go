package config

type Response struct {
	StatusCode int    `yaml:"status_code"`
	Body       string `yaml:"body"`
}
