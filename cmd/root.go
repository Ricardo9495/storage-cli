package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	BaseURL string

	rootCmd = &cobra.Command{
		Use:   "storage-cli",
		Short: "CLI tool for interacting with the storage server",
	}
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found. Using default base URL.")
	}

	// Read from env or fallback
	BaseURL = os.Getenv("STORAGE_API_URL")
	if BaseURL == "" {
		BaseURL = "http://localhost:8080"
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
