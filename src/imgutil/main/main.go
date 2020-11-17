package main

import (
	"fmt"
	"github.com/Jecced/go-tools/src/fileutil"
	"github.com/Jecced/go-tools/src/imgutil"
	"os"
	"path/filepath"
)

func main() {
	var err error
	var dir string

	if len(os.Args) > 1 {
		dir, err = filepath.Abs(os.Args[1])
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if dir == "" {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	list, err := fileutil.GetFilesBySuffix(dir, ".png")
	if err != nil {
		fmt.Println(err)
		return
	}

	c := 0

	for _, path := range list {
		image, err := imgutil.LoadImage(path)
		if err != nil {
			fmt.Printf("=====\n读取错误\n%s\n%s\n", path, err)
			continue
		}

		img, err := imgutil.TrimBlankImg(image)
		if err != nil {
			fmt.Printf("=====\ntrim错误\n%s\n%s\n", path, err)
			continue
		}

		_, name := filepath.Split(path)

		if image.Bounds().Max.X == img.Bounds().Max.X && image.Bounds().Max.Y == img.Bounds().Max.Y {
			fmt.Printf(
				"%7s: %4d -%4d %s 文件尺寸没变化\n",
				"skip",
				image.Bounds().Max.X,
				image.Bounds().Max.Y,
				name,
			)
			continue
		}

		err = imgutil.SaveImage(path, img)
		if err != nil {
			fmt.Printf("=====\n保存错误\n%s\n%s\n", path, err)
			continue
		}
		fmt.Printf(
			"%7s: %4d -%4d  →%4d -%4d %s \n",
			"success",
			image.Bounds().Max.X,
			image.Bounds().Max.Y,
			img.Bounds().Max.X,
			img.Bounds().Max.Y,
			name,
		)
		c++
	}

	fmt.Printf("共计处理文件数量:%d\n", c)

}
