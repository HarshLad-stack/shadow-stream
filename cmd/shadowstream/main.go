package main

import (
	"fmt"
	"io"
	"log"

	"shadow-stream/internal/ffmpeg"
)

func main() {
	reader, err := ffmpeg.StreamRawVideo("test.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	buffer := make([]byte, 4096)
	total := 0

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			total += n
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Total raw bytes read:", total)
}
