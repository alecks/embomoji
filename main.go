package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func main() {
	err := filepath.Walk("in", func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		img, err := imaging.Open(p)
		if err != nil {
			return err
		}
		img = imaging.Convolve3x3(
			img,
			[9]float64{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			nil,
		)

		return imaging.Save(img, path.Join("out", p))
	})
	if err != nil {
		log.Fatalln("Failed:", err)
	}
}
