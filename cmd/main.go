package main

import (
	"fmt"
	"singleton-go/pkg/singleton"
)

func main() {
	fmt.Println("Singleton Pattern in Go")

	// Get the singleton instance of ConfigManager
	config := singleton.GetConfigManager()
	fmt.Println("Database connection string:", config.Get("databaseConnectionString"))

	// Attempt to retrieve the same instance again
	configAgain := singleton.GetConfigManager()
	if config == configAgain {
		fmt.Println("Confirmed that config and configAgain are the same instance.")
	}
}
