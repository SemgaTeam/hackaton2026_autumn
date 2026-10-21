package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the application configuration.
type Config struct {
	Port int
}

// Load reads .env (if present) and environment variables,
// then returns a validated Config.
func Load() (*Config, error) {
	loadDotEnv(".env")

	rawPort := os.Getenv("PORT")
	if rawPort == "" {
		rawPort = "8080"
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT %q: %w", rawPort, err)
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("PORT out of range: %d", port)
	}

	return &Config{Port: port}, nil
}

// Addr returns the address the HTTP server should listen on.
func (c Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}

// loadDotEnv loads KEY=VALUE pairs from path into the process environment.
// Existing environment variables are not overwritten. A missing file is ignored.
func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimPrefix(line, "export ")

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)

		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}
}
