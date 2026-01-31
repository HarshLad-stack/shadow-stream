package ffmpeg

import (
	"io"
	"os/exec"
)

func StreamRawVideo(inputPath string) (io.ReadCloser, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputPath,
		"-f", "rawvideo",
		"-pix_fmt", "rgb24",
		"-",
	) ///---> ffmpeg -i input.mp4 -f rawvideo -pix_fmt rgb24 -
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return stdout, nil

}
