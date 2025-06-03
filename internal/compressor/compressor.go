package compressor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompressImage(inputPath string) {
	ext := strings.ToLower(filepath.Ext(inputPath))
	// For HEIC, output to PNG
	outputExt := ext
	if ext == ".heic" {
		outputExt = ".png"
	}
	// Create a temporary file for compression
	tempFile, err := os.CreateTemp("", "imgpim-*"+outputExt)
	if err != nil {
		fmt.Printf("Error creating temporary file for %s: %v\n", inputPath, err)
		return
	}
	tempPath := tempFile.Name()
	tempFile.Close()

	var cmd *exec.Cmd
	switch ext {
	case ".jpg", ".jpeg":
		// Use mozjpeg's jpegtran for lossless JPEG compression
		cmd = exec.Command("jpegtran", "-optimize", "-progressive", "-copy", "all", "-outfile", tempPath, inputPath)
	case ".png":
		// Use oxipng for lossless PNG compression
		cmd = exec.Command("oxipng", "-o", "max", "--strip", "safe", "--out", tempPath, inputPath)
	case ".gif":
		// Use gifsicle for lossless GIF compression
		cmd = exec.Command("gifsicle", "--optimize=3", "--lossy=0", inputPath, "-o", tempPath)
	case ".heic":
		// Use libheif's heif-convert to PNG for lossless output
		cmd = exec.Command("heif-convert", "--png-compression-level", "9", inputPath, tempPath)
	case ".webp", ".bmp", ".tiff":
		fmt.Printf("Skipping %s: Compression not yet supported\n", inputPath)
		return
	default:
		fmt.Printf("Skipping %s: Unsupported format\n", inputPath)
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error compressing %s: %v\n", inputPath, err)
		os.Remove(tempPath)
		return
	}

	// Replace the input file with the compressed file
	outputPath := strings.TrimSuffix(inputPath, ext) + outputExt
	if err := os.Rename(tempPath, outputPath); err != nil {
		fmt.Printf("Error replacing %s with compressed file: %v\n", inputPath, err)
		os.Remove(tempPath)
		return
	}

	fmt.Printf("Compressed %s\n", inputPath)
}