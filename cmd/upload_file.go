/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var uploadPath string

// uploadFileCmd represents the uploadFile command
var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload a file to the storage server",
	Run: func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			fmt.Println("Please provide --name flag")
			return
		}

		if uploadPath == "" {
			fmt.Println("Please provide --file flag")
			return
		}

		file, err := os.Open(uploadPath)
		if err != nil {
			fmt.Println("Failed to open file:", err)
			return
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", file.Name())
		if err != nil {
			fmt.Println("Failed to create form file:", err)
			return
		}
		io.Copy(part, file)
		writer.Close()

		resp, err := http.Post(BaseURL+"/"+fileName, writer.FormDataContentType(), body)
		if err != nil {
			fmt.Println("Upload failed:", err)
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
				fmt.Printf("Upload failed: %s - %s\n", resp.Status, result["error"])
			} else {
				fmt.Printf("Upload failed: %s - %s\n", resp.Status, string(bodyBytes))
			}
			return
		}

		fmt.Println("Upload response:", resp.Status)
	},
}

func init() {
	uploadFileCmd.Flags().StringVar(&fileName, "name", "", "Name of the file")
	uploadFileCmd.Flags().StringVar(&uploadPath, "file", "", "Path to the file")
	rootCmd.AddCommand(uploadFileCmd)
}
