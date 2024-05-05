package singleton

import (
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
		// Assume settings are loaded here
		instance.settings["databaseConnectionString"] = "Server=myServerAddress;Database=myDataBase;Uid=myUsername;Pwd=myPassword;"
	})
	return instance
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
