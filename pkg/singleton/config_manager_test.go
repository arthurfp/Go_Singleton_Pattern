package singleton

import (
	"sync"
	"testing"
)

// TestSingletonInstance tests whether multiple calls to GetConfigManager return the same instance.
func TestSingletonInstance(t *testing.T) {
	instance1 := GetConfigManager()
	instance2 := GetConfigManager()

	if instance1 != instance2 {
		t.Errorf("GetConfigManager did not return the same instance")
	}
}

// TestSingletonThreadSafety tests the thread safety of setting and getting configuration values.
func TestSingletonThreadSafety(t *testing.T) {
	config := GetConfigManager()
	var wg sync.WaitGroup
	iterations := 100

	// Test concurrent access to Set method
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			configKey := "testKey"
			configValue := "testValue" + string(val)
			config.Set(configKey, configValue)

			if gotValue := config.Get(configKey); gotValue != configValue {
				t.Errorf("Got incorrect value, expected %s, got %s", configValue, gotValue)
			}
		}(i)
	}
	wg.Wait()
}
