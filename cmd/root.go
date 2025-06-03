package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/Ansh-Rathod/imgpim/internal/compressor"
	"github.com/Ansh-Rathod/imgpim/internal/utils"
)

var rootCmd = &cobra.Command{
	Use:   "imgpim [input]",
	Short: "Compress images losslessly using open-source tools",
	Long:  `imgpim is a CLI tool to compress images (jpg, jpeg, png, gif, heic) losslessly in place using tools like oxipng, gifsicle, mozjpeg, and libheif.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Printf("Error: Input path does not exist: %s\n", inputPath)
			os.Exit(1)
		}

		// Check dependencies
		tools := []string{"oxipng", "gifsicle", "jpegtran", "heif-convert"}
		if missing := utils.CheckDependencies(tools); len(missing) > 0 {
			fmt.Printf("Error: Missing dependencies: %v\n", missing)
			fmt.Println("Please install them using Homebrew or your package manager.")
			os.Exit(1)
		}

		// Process input
		if info, _ := os.Stat(inputPath); info.IsDir() {
			err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && utils.IsSupportedImage(path) {
					compressor.CompressImage(path)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error processing directory: %v\n", err)
				os.Exit(1)
			}
		} else {
			if utils.IsSupportedImage(inputPath) {
				compressor.CompressImage(inputPath)
			} else {
				fmt.Printf("Skipping %s: Unsupported format\n", inputPath)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}