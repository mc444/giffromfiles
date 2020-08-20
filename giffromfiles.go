//Giffromfiles generates an animated GIF from source GIF files.
// &image.Paletted{
//   Pix:[]uint8{...},
//   Stride: 800,
//   Rect: image.Rectangle{Min:image.Point{X:0, Y:0}, Max:image.Point{X:800, Y:600}}
//   Palette: color.Palette{
//
package main

import (
	"image"
	"image/gif"
	"log"
	"os"
	"path/filepath"
)

func main() {
	const (
		delay = 10 //delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: -1}
	filenames, err := filepath.Glob("./runs/*.gif")
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range filenames {
		reader, err := os.Open("./" + filename)
		if err != nil {
			log.Fatal(err)
		}
		data, err := gif.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		reader.Close()
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, data.(*image.Paletted))
	}
	err = gif.EncodeAll(os.Stdout, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
