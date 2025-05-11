package configs

import (
	"os"
	"strconv"
)

// Config holds the server configuration
type Config struct {
	// Port is the port number the server will listen on
	Port int
	// TerminationGracePeriodInSeconds is the time in seconds to wait for graceful shutdown
	TerminationGracePeriodInSeconds int
}

// New creates a new Config with default values
func New() *Config {
	return &Config{
		Port:                            8080,
		TerminationGracePeriodInSeconds: 5,
	}
}

// LoadFromEnv loads configuration from environment variables
func (c *Config) LoadFromEnv() {
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			c.Port = p
		}
	}

	if gracePeriod := os.Getenv("TERMINATION_GRACE_PERIOD_IN_SECONDS"); gracePeriod != "" {
		if gp, err := strconv.Atoi(gracePeriod); err == nil {
			c.TerminationGracePeriodInSeconds = gp
		}
	}
}
