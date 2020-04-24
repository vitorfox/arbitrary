package config

type Route struct {
	Path               string     `yaml:"path"`
	Method             string     `yaml:"method"`
	SuccessfulResponse Response   `yaml:"successful_response"`
	ThrottledResponse  Response   `yaml:"throttled_response"`
}
