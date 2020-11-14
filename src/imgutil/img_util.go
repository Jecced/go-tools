package imgutil

import (
	"image"
	"image/draw"
)

// 图片裁剪
func Trimming(in, out string, x, y, w, h int) error {
	src, err := LoadImage(in)
	if err != nil {
		return err
	}

	img, err := imageCopy(src, x, y, w, h)
	if err != nil {
		return err
	}
	err = SaveImage(out, img)
	if err != nil {
		return err
	}
	return nil
}

// 创建一个指定宽高的图片
func CreatPng(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// 混合两个图片
func MixImg(src *image.RGBA, dist image.Image, x, y int) {
	draw.Draw(src, dist.Bounds().Add(image.Pt(x, y)), dist, image.Point{}, draw.Over)
}
