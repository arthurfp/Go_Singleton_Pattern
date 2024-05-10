# Singleton Pattern Demonstration in Go

## Overview
This repository demonstrates the implementation of the Singleton pattern in Go. The project focuses on creating a globally accessible configuration manager (`ConfigManager`) that maintains a single instance throughout the application's lifecycle. This ensures that configuration settings are managed centrally and efficiently, minimizing resource use and ensuring consistency.

## Pattern Description
The Singleton Pattern ensures that a class has only one instance and provides a global point of access to it. This is particularly useful in managing shared resources such as configuration settings, where a single point of control is necessary to manage state across an application. In this project, the `ConfigManager` class is designed as a Singleton to handle configuration settings, making it easy to manage and update settings dynamically.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the use of the Singleton pattern with the `ConfigManager`.
- **pkg/**
  - **singleton/**: Implements the `ConfigManager`, the Singleton class managing the application's configuration.
- **internal/**
  - **database/**: Contains the simulated `Connection` class used for demonstrating the Singleton's use in managing database connections.

## Configuration
Environment variables are used to initialize the settings in the `ConfigManager`. The following variables can be set:
- `DATABASE_CONNECTION_STRING`: Database connection string.
- `API_KEY`: API key for external services.
- `LOG_LEVEL`: Logging level for the application.

## Features
- **Environment Variable Support**: Allows initial configuration through environment variables.
- **Dynamic Setting Updates**: Supports updating settings at runtime through the Singleton interface.
- **Detailed Logging**: Logs all interactions with the configuration settings for auditing and debugging purposes.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Singleton_Pattern.git
cd Go_Singleton_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp
