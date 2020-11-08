package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"syscall/js"

	"github.com/EdlinOrg/prominentcolor"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")

	js.Global().Set("findCommonColors", js.FuncOf(findCommonColors))

	<-c
}

func findCommonColors(this js.Value, args []js.Value) interface{} {
	sourceImage := getImageFromArgs(args)
	getDominantColors(sourceImage)
	colors := getDominantColors(sourceImage)

	return js.ValueOf(colors)
}

func getImageFromArgs(args []js.Value) image.Image {
	array := args[0]
	inBuf := make([]uint8, array.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, array)

	reader := bytes.NewReader(inBuf)

	sourceImage, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	return sourceImage
}

func getDominantColors(image image.Image) []interface{} {
	colors, err := prominentcolor.Kmeans(image)

	if err != nil {
		log.Fatal("Failed to get dominant colors", err)
	}

	colorsHex := make([]interface{}, 3)

	for i, color := range colors {
		colorsHex[i] = "#" + color.AsString()
	}

	return colorsHex
}
