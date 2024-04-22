package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // Import to support decoding PNG images.
	"os"
	"strconv"
	"strings"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

func main() {
	// Command-line flags setup.
	sourcePath := flag.String("source", "", "Path to the source image file")
	targetPath := flag.String("target", "", "Path to the target (resized) image file")
	widths := flag.String("widths", "", "List of widths to resize to, separated by commas")
	format := flag.String("format", "webp", "Format of the target image (webp or jpg)")

	flag.Parse()

	// Validate required flags.
	if *sourcePath == "" || *targetPath == "" || *widths == "" {
		fmt.Println("Missing required arguments.")
		fmt.Println("Usage: go run main.go -source=path/to/source.jpg -target=path/to/target -widths=1080,720,320 -format=webp")
		os.Exit(1)
	}

	// Validate format.
	if *format != "jpg" && *format != "webp" {
		fmt.Println("Invalid format specified. Supported formats: 'webp', 'jpg'.")
		os.Exit(1)
	}

	// Open source image file.
	file, err := os.Open(*sourcePath)
	if err != nil {
		fmt.Printf("Error opening source image: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode image from file.
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		os.Exit(1)
	}

	// Process each width specified.
	for _, w := range strings.Split(*widths, ",") {
		width, err := strconv.Atoi(w)
		if err != nil {
			fmt.Printf("Invalid width '%s': %v\n", w, err)
			os.Exit(1)
		}

		// Resize the image.
		resizedImage := resizeImage(img, width)

		// Prepare the output filename.
		outputFileName := fmt.Sprintf("%s-%d.%s", *targetPath, width, *format)

		// Create output file.
		out, err := os.Create(outputFileName)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			continue
		}

		defer out.Close()

		// Encode and save the resized image.
		switch *format {
		case "jpg":
			err := jpeg.Encode(out, resizedImage, &jpeg.Options{Quality: 90})
			if err != nil {
				fmt.Printf("Error encoding JPEG image: %v\n", err)
				continue
			}
		case "webp":
			err := webp.Encode(out, resizedImage, &webp.Options{Quality: 90})
			if err != nil {
				fmt.Printf("Error encoding WEBP image: %v\n", err)
				continue
			}
		}

		fmt.Printf("Resized image saved to %s\n", outputFileName)
	}
}

// resizeImage resizes an image to the specified width while maintaining the aspect ratio.
func resizeImage(img image.Image, width int) image.Image {
	srcBounds := img.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()
	height := (width * srcHeight) / srcWidth // Calculate the new height based on aspect ratio.

	newRect := image.Rect(0, 0, width, height)
	newImg := image.NewRGBA(newRect)

	// Use approximated bi-linear scaling to resize the image.
	draw.ApproxBiLinear.Scale(newImg, newRect, img, srcBounds, draw.Over, nil)

	return newImg
}
