package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/mattn/go-sixel"
)

func main() {
	file, _ := os.Open("./gopher.png")
	image, _, _ := image.Decode(file)
	buf := bytes.Buffer{}
	_ = sixel.NewEncoder(&buf).Encode(image)
	fmt.Println(buf.String())

	out, _ := os.Create("gopher.sixel")
	_, _ = out.Write(buf.Bytes())
}
