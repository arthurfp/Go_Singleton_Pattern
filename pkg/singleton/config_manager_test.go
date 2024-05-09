package singleton

import (
	"sync"
	"testing"
)

// TestSingletonInstance ensures that multiple calls to GetConfigManager return the same instance.
func TestSingletonInstance(t *testing.T) {
	manager1 := GetConfigManager()
	manager2 := GetConfigManager()

	if manager1 != manager2 {
		t.Errorf("GetConfigManager did not return the same instance")
	}
}

// TestSingletonThreadSafety tests the thread safety of the Singleton implementation.
func TestSingletonThreadSafety(t *testing.T) {
	var wg sync.WaitGroup
	iterations := 100
	results := make([]*ConfigManager, iterations)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			results[index] = GetConfigManager()
		}(i)
	}
	wg.Wait()

	// Check that all instances are the same
	for _, instance := range results {
		if instance != results[0] {
			t.Errorf("Singleton instance is not the same across goroutines")
		}
	}
}

// TestConfigSettings tests setting and getting configuration values.
func TestConfigSettings(t *testing.T) {
	manager := GetConfigManager()
	testKey := "testKey"
	testValue := "testValue"

	manager.Set(testKey, testValue)
	retrievedValue := manager.Get(testKey)
	if retrievedValue != testValue {
		t.Errorf("Expected value %s, got %s", testValue, retrievedValue)
	}
}
