package compressor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompressImage(inputPath, output string, quality int) {
	ext := strings.ToLower(filepath.Ext(inputPath))
	outputPath := output
	if output == "" {
		outputPath = inputPath
	} else if info, _ := os.Stat(output); info != nil && info.IsDir() {
		outputPath = filepath.Join(output, filepath.Base(inputPath))
	}

	var cmd *exec.Cmd
	switch ext {
	case ".jpg", ".jpeg":
		// Use mozjpeg's jpegtran for JPEG
		cmd = exec.Command("jpegtran", "-optimize", "-progressive", "-outfile", outputPath, inputPath)
	case ".png":
		// Use oxipng for PNG
		cmd = exec.Command("oxipng", "-o", "max", "--strip", "safe", inputPath, "-o", outputPath)
	case ".gif":
		// Use gifsicle for GIF
		cmd = exec.Command("gifsicle", "--optimize=3", inputPath, "-o", outputPath)
	case ".heic":
		// Use libheif's heif-convert
		cmd = exec.Command("heif-convert", inputPath, outputPath)
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
		return
	}
	fmt.Printf("Compressed %s -> %s\n", inputPath, outputPath)
}