package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func main() {
	if _, err := os.Stat("out"); err != nil {
		os.Mkdir("out", os.ModePerm)
	}
	err := filepath.Walk("in", func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
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

		_, filename := path.Split(p)
		return imaging.Save(img, path.Join("out", filename))
	})
	if err != nil {
		log.Fatalln("Failed:", err)
	}
}
