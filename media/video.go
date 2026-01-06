package media

import "os/exec"

func ProcessVideo(inputPath string, outputPath string) error {
	// Placeholder for video processing logic
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputPath,
		"-vf", "scale=100:100:flags=lanczos,fps=30",
		"-t", "3",
		"-an",
		"-c:v", "libvpx-vp9",
		"-pix_fmt", "yuva420p",
		"-b:v", "0",
		"-crf", "36",
		"-loop", "1",
		outputPath,
	)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
