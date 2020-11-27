package imgutil

import (
	"github.com/Jecced/go-tools/src/fileutil"
	"image"
)

// 将一个文件夹内的内容都进行Trim
func TrimBlankDir(dir, suffix string) error {
	list, err := fileutil.GetFilesBySuffix(dir, suffix)
	if err != nil {
		return err
	}
	for _, path := range list {
		err = TrimBlankFile(path, path)
		if err != nil {
			return err
		}
	}
	return nil
}

// 将一个图片文件Trim空白边
func TrimBlankFile(in, out string) error {
	src, err := LoadImage(in)
	if err != nil {
		return err
	}

	img, err := TrimBlankImg(src)
	if err != nil {
		return err
	}
	err = SaveImage(out, img)
	if err != nil {
		return err
	}

	return nil
}

// 将一个图片对象Trim空白边
func TrimBlankImg(src image.Image) (image.Image, error) {
	bounds := src.Bounds()
	minX, minY, maxX, maxY := bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

	for x, y := minX, minY; y < maxY; x++ {
		if x > maxX {
			y++
			x = 0
			continue
		}
		_, _, _, a := src.At(x, y).RGBA()
		if a != 0 {
			minY = y
			break
		}
	}

	for x, y := minX, maxY; y >= minY; x++ {
		if x > maxX {
			y--
			x = 0
			continue
		}
		_, _, _, a := src.At(x, y).RGBA()
		if a != 0 {
			maxY = y + 1
			break
		}
	}

	for x, y := minX, minY; x < maxX; y++ {
		if y > maxY {
			x++
			y = minY
			continue
		}
		_, _, _, a := src.At(x, y).RGBA()
		if a != 0 {
			minX = x
			break
		}
	}

	for x, y := maxX, minY; x >= minX; y++ {
		if y > maxY {
			x--
			y = minY
			continue
		}
		_, _, _, a := src.At(x, y).RGBA()
		if a != 0 {
			maxX = x + 1
			break
		}
	}

	return imageCopy(src, minX, minY, maxX-minX, maxY-minY)
}
