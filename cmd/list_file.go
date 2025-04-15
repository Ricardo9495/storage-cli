/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"mymodule/cmd/entity"
	"net/http"

	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFileCmd = &cobra.Command{
	Use:   "list-file",
	Short: "List all files from the server",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(BaseURL)
		if err != nil {
			fmt.Println("Request failed:", err)
			return
		}
		defer resp.Body.Close()

		var files []entity.File
		err = json.NewDecoder(resp.Body).Decode(&files) // Decode directly into []entity.File
		if err != nil {
			fmt.Println("Failed to decode response:", err)
			return
		}

		for _, file := range files {
			jsonData, err := json.Marshal(file)
			if err != nil {
				fmt.Println("Failed to marshal file to JSON:", err)
				continue
			}
			fmt.Println(string(jsonData)) // Print each JSON object on a new line
		}
	},
}

func init() {
	rootCmd.AddCommand(listFileCmd)
}
