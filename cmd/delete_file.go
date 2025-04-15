/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var fileName string

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Delete a file by filename",
	Run: func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			fmt.Println("Please provide --name flag")
			return
		}
		req, err := http.NewRequest("DELETE", BaseURL+"/"+fileName, nil)
		if err != nil {
			fmt.Println("Request error:", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Delete failed:", err)
			return
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			var result map[string]interface{}
			err = json.Unmarshal(bodyBytes, &result)
			if err == nil && result["error"] != nil {
				fmt.Printf("Delete failed: %s - %s\n", resp.Status, result["error"])
			} else {
				fmt.Printf("Delete failed: %s - %s\n", resp.Status, string(bodyBytes))
			}
			return
		}

		fmt.Println("Delete response:", resp.Status)
	},
}

func init() {
	deleteFileCmd.Flags().StringVar(&fileName, "name", "", "Name of the file to delete")
	deleteFileCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(deleteFileCmd)
}
