package utils

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func CheckDependencies(tools []string) []string {
	var missing []string
	for _, tool := range tools {
		if _, err := exec.LookPath(tool); err != nil {
			missing = append(missing, tool)
		}
	}
	return missing
}

func IsSupportedImage(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".heic" ||
		ext == ".webp" || ext == ".bmp" || ext == ".tiff"
}