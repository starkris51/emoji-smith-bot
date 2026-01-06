package media

import (
	"os/exec"
)

func ProcessImage(inputPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=100:100", outputPath)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
