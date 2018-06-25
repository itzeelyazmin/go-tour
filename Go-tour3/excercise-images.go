package main

import ("golang.org/x/tour/pic"
"image"
"image/color"
)

type Image struct{
width int
height int}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.width, i.height)
}


func main() {
	i := Image{256, 256}
	pic.ShowImage(i)
}