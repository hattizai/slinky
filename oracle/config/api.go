package config

import (
	"fmt"
	"time"
)

// APIConfig defines a config for an API based data provider.
type APIConfig struct {
	// Enabled is a flag that indicates whether the provider is API based.
	Enabled bool `mapstructure:"enabled" toml:"enabled"`

	// Timeout is the amount of time the provider should wait for a response from
	// its API before timing out.
	Timeout time.Duration `mapstructure:"timeout" toml:"timeout"`

	// Interval is the interval at which the provider should update the prices.
	Interval time.Duration `mapstructure:"interval" toml:"interval"`

	// MaxQueries is the maximum number of queries that the provider will make
	// within the interval. If the provider makes more queries than this, it will
	// stop making queries until the next interval.
	MaxQueries int `mapstructure:"max_queries" toml:"max_queries"`
}

func (c *APIConfig) ValidateBasic() error {
	if !c.Enabled {
		return nil
	}

	if c.MaxQueries < 1 {
		return fmt.Errorf("api max queries must be greater than 0")
	}

	if c.Interval <= 0 || c.Timeout <= 0 {
		return fmt.Errorf("provider interval and timeout must be strictly positive")
	}

	if c.Interval < c.Timeout {
		return fmt.Errorf("provider timeout must be greater than 0 and less than the interval")
	}

	return nil
}
