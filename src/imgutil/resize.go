package imgutil

import (
	"image"
	"image/color"
)

// https://studygolang.com/articles/1156
// Resample返回图像切片r (m)的重新采样副本。
// 返回的图像宽度为w，高度为h。
//func Resample(m image.Image, r image.Rectangle, w, h int) *image.RGBA {
func Resize(m image.Image, w, h int) *image.RGBA {
	if w < 0 || h < 0 {
		return nil
	}
	r := m.Bounds()
	curw, curh := r.Dx(), r.Dy()
	if w == 0 || h == 0 || curw <= 0 || curh <= 0 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Get a source pixel.
			subx := x * curw / w
			suby := y * curh / h
			r32, g32, b32, a32 := m.At(subx, suby).RGBA()
			r := uint8(r32 >> 8)
			g := uint8(g32 >> 8)
			b := uint8(b32 >> 8)
			a := uint8(a32 >> 8)
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}
	return img
}