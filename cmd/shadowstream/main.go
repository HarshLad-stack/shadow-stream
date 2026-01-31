package main

import (
	"io"
	"log"

	"shadow-stream/internal/ffmpeg"
	"shadow-stream/internal/steg"
)

func main() {
	reader, err := ffmpeg.StreamRawVideo("test.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// buffer size MUST be multiple of 3 (RGB)
	buffer := make([]byte, 4096*3)

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			// operate only on valid bytes
			steg.EmbedBitsInStream(buffer[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
