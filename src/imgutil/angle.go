package imgutil

import "image"

// img 图像, 顺时针旋转90°
func RotationRight(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	rotate90 := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
	for x := bounds.Min.Y; x < bounds.Max.Y; x++ {
		for y := bounds.Max.X - 1; y >= bounds.Min.X; y-- {
			rotate90.Set(bounds.Max.Y-x, y, m.At(y, x))
		}
	}
	return rotate90
}

// img 图像, 逆时针旋转90°
func RotationLeft(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	rotate270 := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
	for x := bounds.Min.Y; x < bounds.Max.Y; x++ {
		for y := bounds.Max.X - 1; y >= bounds.Min.X; y-- {
			rotate270.Set(x, bounds.Max.X-y, m.At(y, x))
		}
	}
	return rotate270
}

// img 图像, 180°旋转
func Rotation180(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	rotate180 := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rotate180.Set(bounds.Max.X-x, bounds.Max.Y-y, m.At(x, y))
		}
	}
	return rotate180
}

// img 图像, 左右镜像翻转
func FlipMirror(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			out.Set(bounds.Max.X-x, y, m.At(x, y))
		}
	}
	return out
}

// img 图像, 上下垂直翻转
func FlipVertical(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			out.Set(x, bounds.Max.Y-y, m.At(x, y))
		}
	}
	return out
}

// img 文件, 顺时针旋转90°
func RotationImgRight(in, out string) error {
	return rotation(in, out, RotationRight)
}

// img 文件, 逆时针旋转90°
func RotationImgLeft(in, out string) error {
	return rotation(in, out, RotationLeft)
}

// img 文件, 180°旋转
func RotationImg180(in, out string) error {
	return rotation(in, out, Rotation180)
}

// img 文件 镜像翻转
func FlipImgMirror(in, out string) error {
	return rotation(in, out, FlipMirror)
}

// img 文件 上下垂直翻转
func FlipImgVertical(in, out string) error {
	return rotation(in, out, FlipVertical)
}

func rotation(in, out string, meth func(image.Image) *image.RGBA) error {
	loadImage, err := LoadImage(in)
	if err != nil {
		return err
	}
	img := meth(loadImage)
	err = SaveImage(out, img)
	if err != nil {
		return err
	}
	return nil
}
