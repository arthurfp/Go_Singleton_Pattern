package singleton

import (
	"log"
	"os"
	"sync"
)

// ConfigManager manages application configuration settings in a thread-safe manner.
// It uses the Singleton design pattern to ensure only one instance manages the configuration settings across the application.
type ConfigManager struct {
	sync.RWMutex
	settings map[string]string
}

var instance *ConfigManager
var once sync.Once

// GetConfigManager returns the singleton instance of the ConfigManager.
// It ensures that the ConfigManager is initialized only once.
func GetConfigManager() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{settings: make(map[string]string)}
		log.Println("Initializing ConfigManager singleton instance.")
		instance.loadEnvironmentVariables()
	})
	return instance
}

// loadEnvironmentVariables initializes configuration settings from environment variables.
// This method logs the success or failure of loading each variable.
func (c *ConfigManager) loadEnvironmentVariables() {
	envVariables := []string{"DATABASE_CONNECTION_STRING", "API_KEY", "LOG_LEVEL"}
	for _, envVar := range envVariables {
		value, exists := os.LookupEnv(envVar)
		if exists {
			c.settings[envVar] = value
			log.Printf("Loaded %s from environment variables.", envVar)
		} else {
			log.Printf("Environment variable %s is not set.", envVar)
		}
	}
}

// Get retrieves a setting value by key and logs the access.
// It is thread-safe.
func (c *ConfigManager) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	log.Printf("Accessed setting %s.", key)
	return c.settings[key]
}

// Set updates a setting value by key and logs the update.
// It ensures thread-safe updates to configuration settings.
func (c *ConfigManager) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.settings[key] = value
	log.Printf("Updated setting %s to %s.", key, value)
}
