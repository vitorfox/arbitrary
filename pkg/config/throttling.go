package config

import (
	"time"
)

type Throttling struct {
	MaxSimultaneousRequests  int           `yaml:"max_simultaneous_requests"`
	MaxThrottledRequests     int           `yaml:"max_throttled_requests"`
	DelayOnResponse          time.Duration `yaml:"delay_on_response"`
	DelayOnThrottledResponse time.Duration `yaml:"delay_on_throttled_response"`
}
