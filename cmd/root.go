package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/Ansh-Rathod/imgpim/internal/compressor"
	"github.com/Ansh-Rathod/imgpim/internal/utils"
)

var output string
var quality int

var rootCmd = &cobra.Command{
	Use:   "imgpim [input]",
	Short: "Compress images using open-source tools",
	Long:  `imgpim is a CLI tool to compress images (jpg, jpeg, png, gif, heic) using tools like oxipng, gifsicle, mozjpeg, jpegoptim, and libheif.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			fmt.Printf("Error: Input path does not exist: %s\n", inputPath)
			os.Exit(1)
		}

		// Check dependencies
		tools := []string{"oxipng", "gifsicle", "jpegtran", "jpegoptim", "heif-convert"}
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
					compressor.CompressImage(path, output, quality)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error processing directory: %v\n", err)
				os.Exit(1)
			}
		} else {
			if utils.IsSupportedImage(inputPath) {
				compressor.CompressImage(inputPath, output, quality)
			} else {
				fmt.Printf("Skipping %s: Unsupported format\n", inputPath)
			}
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output directory or file path")
	rootCmd.Flags().IntVarP(&quality, "quality", "q", 85, "Compression quality (1-100, default 85)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}