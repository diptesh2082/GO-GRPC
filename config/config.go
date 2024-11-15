package config

import "os"

func LoadConfig() {
	// Load configuration from environment variables
	os.Getenv("BILLING_API_PORT")
}