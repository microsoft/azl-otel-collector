// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package smartdata

import (
	"fmt"
	"time"
)

type Config struct {
	Interval string `mapstructure:"interval"`
}

// Validate if the configuration is valid
func (c *Config) Validate() error {
	interval, _ := time.ParseDuration(c.Interval)
	if interval.Minutes() < 1 {
		return fmt.Errorf("interval must be at least 1 minute")
	}
	return nil
}
