package config

type Base struct {
	Routes     []Route    `yaml:"routes"`
	Throttling Throttling `yaml:"throttling"`
}
