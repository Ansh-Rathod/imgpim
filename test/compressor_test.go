package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Ansh-Rathod/imgpim/internal/compressor"
	"github.com/Ansh-Rathod/imgpim/internal/utils"
)

func TestCompressImage(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()
	inputPath := filepath.Join(tempDir, "test.jpg")
	outputPath := filepath.Join(tempDir, "output.jpg")

	// Create a dummy JPEG file (minimal valid JPEG)
	dummyJPEG := []byte{
		0xff, 0xd8, // SOI
		0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46, 0x00, 0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, // APP0
		0xff, 0xd9, // EOI
	}
	if err := os.WriteFile(inputPath, dummyJPEG, 0644); err != nil {
		t.Fatalf("Failed to create dummy JPEG: %v", err)
	}

	// Skip test if jpegtran is not installed
	if missing := utils.CheckDependencies([]string{"jpegtran"}); len(missing) > 0 {
		t.Skip("jpegtran not installed, skipping test")
	}

	// Test compression
	compressor.CompressImage(inputPath, outputPath, 85)

	// Verify output exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Output file was not created")
	}
}