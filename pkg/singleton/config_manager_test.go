package singleton

import (
	"log"
	"os"
	"sync"
)

// ConfigManager manages application configuration settings.
type ConfigManager struct {
	sync.RWMutex
	settings map[string]string
}

var instance *ConfigManager
var once sync.Once

// GetConfigManager returns the singleton instance of the ConfigManager.
func GetConfigManager() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			settings: make(map[string]string),
		}
		// Initialize settings from environment variables
		instance.loadEnvironmentVariables()
	})
	return instance
}

// loadEnvironmentVariables loads configuration settings from environment variables.
func (c *ConfigManager) loadEnvironmentVariables() {
	envVariables := []string{"DATABASE_CONNECTION_STRING", "API_KEY", "LOG_LEVEL"}
	for _, envVar := range envVariables {
		value, exists := os.LookupEnv(envVar)
		if exists {
			c.settings[envVar] = value
		} else {
			log.Printf("Environment variable %s is not set", envVar)
		}
	}
}

// Get retrieves a setting value by key.
func (c *ConfigManager) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	return c.settings[key]
}

// Set updates a setting value by key.
func (c *ConfigManager) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.settings[key] = value
}
