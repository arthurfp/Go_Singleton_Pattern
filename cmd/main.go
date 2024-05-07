package main

import (
	"fmt"
	"os"
	"singleton-go/pkg/singleton"
)

func main() {
	// Set an example environment variable before getting the ConfigManager
	os.Setenv("DATABASE_CONNECTION_STRING", "Server=myServerAddress;Database=myDataBase;Uid=myUsername;Pwd=myPassword;")

	fmt.Println("Singleton Pattern in Go")

	config := singleton.GetConfigManager()
	connectionString := config.Get("DATABASE_CONNECTION_STRING")
	fmt.Println("Database connection string from environment:", connectionString)
}
