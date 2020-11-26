package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/fileutil"
	"path/filepath"
	"testing"
)

func TestRelativePath(t *testing.T) {
	outJsPath := "/Users/ankang/git/saisheng/slgrpg/temp/quick-scripts/assets/script/feature/battleoverride"
	filePath := "/Users/ankang/git/saisheng/slgrpg/assets/script/feature/battleoverride"
	fmt.Println(outJsPath, filePath)
	path := fileutil.GetRelativePath(outJsPath, filePath)
	fmt.Println(path)

	rel, err := filepath.Rel(outJsPath, filePath)
	if err != nil {
		return
	}
	fmt.Println(222)
	fmt.Println(rel)
}
