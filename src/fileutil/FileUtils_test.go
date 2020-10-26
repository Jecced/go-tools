package fileutil

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	outJsPath := "/Users/ankang/git/saisheng/slgrpg/temp/quick-scripts/assets/script/feature/battleoverride"
	filePath := "/Users/ankang/git/saisheng/slgrpg/assets/script/feature/battleoverride"
	fmt.Println(outJsPath, filePath)
	path := GetRelativePath(outJsPath, filePath)
	fmt.Println(path)

	rel, err := filepath.Rel(outJsPath, filePath)
	if err != nil {
		return
	}
	fmt.Println(222)
	fmt.Println(rel)
}
