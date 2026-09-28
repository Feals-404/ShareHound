// Package config provides configuration management for ShareHound.
package config

import (
	"runtime"
)

// Config holds the configuration settings for ShareHound.
type Config struct {
	debug      bool
	noColors   bool
	smbDialect uint16 // forced SMB dialect (0 = auto-negotiate)
}

// NewConfig creates a new Config with the given settings.
// If noColors is nil, it defaults based on the platform.
func NewConfig(debug bool, noColors *bool) *Config {
	cfg := &Config{
		debug: debug,
	}

	if noColors != nil {
		cfg.noColors = *noColors
	} else {
		// Platform-specific default: disable colors on non-Linux by default
		if runtime.GOOS != "linux" {
			cfg.noColors = true
		} else {
			cfg.noColors = false
		}
	}

	return cfg
}

// Debug returns whether debug mode is enabled.
func (c *Config) Debug() bool {
	return c.debug
}

// SetDebug sets the debug mode.
func (c *Config) SetDebug(value bool) {
	c.debug = value
}

// NoColors returns whether colored output is disabled.
func (c *Config) NoColors() bool {
	return c.noColors
}

// SetNoColors sets whether colored output is disabled.
func (c *Config) SetNoColors(value bool) {
	c.noColors = value
}

// SMBDialect returns the forced SMB dialect (0 = auto-negotiate).
func (c *Config) SMBDialect() uint16 {
	return c.smbDialect
}

// SetSMBDialect sets the forced SMB dialect (0 = auto-negotiate).
func (c *Config) SetSMBDialect(value uint16) {
	c.smbDialect = value
}
