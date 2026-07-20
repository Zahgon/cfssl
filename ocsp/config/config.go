package config

import "time"

type Config struct {
	CACertFile        string
	ResponderCertFile string
	KeyFile           string
	Interval          time.Duration
}
