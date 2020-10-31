package fileutil

import (
	"bytes"
	"fmt"
	"github.com/Jecced/go-tools/src/ak"
	"github.com/Jecced/go-tools/src/commutil"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// 系统路径分隔符
	FileSep = ak.PS
)

// 根据路径创建文件夹
func MkdirAll(path string) {
	_ = os.MkdirAll(path, 0777)
}

// 创建一个文件的父目录
func MkdirParent(path string) {
	parent := GetParentDir(path)
	if !PathExists(parent) {
		MkdirAll(parent)
	}
}

// 获取某个文件夹下所有指定后缀的文件
func GetFilesBySuffix(dirPath string, suffix string) (files []string, err error) {
	// 后缀转大写
	suffix = strings.ToUpper(suffix)
	for strings.HasSuffix(dirPath, FileSep) {
		dirPath = dirPath[:len(dirPath)-1]
	}
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		if fi.IsDir() {
			newFiles, _ := GetFilesBySuffix(dirPath+FileSep+fi.Name(), suffix)
			files = append(files, newFiles...)
		} else if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPath+FileSep+fi.Name())
		}
	}
	return files, nil
}

// 文件拷贝
func FileCopy(src, dist string) (length int, err error) {
	_ = os.Remove(dist)

	// 开启 源文件
	srcFile, err := os.Open(src)
	if err != nil {
		_ = fmt.Errorf("src file open faild, name: %s, error: %v\n", src, err)
		return 0, err
	}
	defer srcFile.Close()

	// 创建输出文件的父目录
	MkdirParent(dist)

	// 创建目标文件
	distFile, err := os.Create(dist)
	if err != nil {
		_ = fmt.Errorf("dist file creat faild, name: %s, error: %v\n", src, err)
		return 0, err
	}
	defer distFile.Close()

	// 创建缓冲区
	bs := make([]byte, 1024*10, 1024*10)
	n := -1
	total := 0
	for {
		n, err = srcFile.Read(bs)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return total, err
		}
		total += n
		_, _ = distFile.Write(bs[:n])
	}
	return total, nil
}

// 判断一个路径是否存在
func PathExists(path string) bool {
	stat, _ := os.Stat(path)
	return stat != nil
}

// 获取一个路径的父目录地址
func GetParentDir(path string) string {
	//path = strings.Trim(path, " ")
	//if strings.HasSuffix(path, "/") || strings.HasSuffix(path, FileSep) {
	//	path = path[0 : len(path)-1]
	//}
	//index := strings.LastIndex(path, "/")
	//if -1 == index {
	//	index = strings.LastIndex(path, FileSep)
	//}
	//return path[0:index]
	dir, _ := filepath.Split(path)
	return dir
}

// 路径格式化, 标准化一个路径到当前系统规范
func PathFormat(path string) string {
	list := []rune(path)
	ps := os.PathSeparator
	for i, l, t := 0, len(list), ' '; i < l; i++ {
		t = list[i]
		// 不是分隔符的字节, 直接跳过
		if t != '/' && t != '\\' {
			continue
		}
		// 将当前字节先替换成系统分隔符
		list[i] = ps
		//如果当前字符和上一个字符相同, 则删除, 否则, 跳过
		if i == 0 || list[i] != list[i-1] {
			continue
		}
		list = append(list[:i], list[i+1:]...)
		i--
		l--
	}
	return string(list)
}

// 获取目录下所有文件类型
func FindAllFileTypes(dir string) (types []string) {

	cache := make(map[string]bool)

	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	var name string
	var i int
	for _, info := range readDir {
		if info.IsDir() {
			out := FindAllFileTypes(dir + ak.PS + info.Name())
			for _, s := range out {
				cache[s] = true
			}
			continue
		}
		name = info.Name()
		name = strings.ToLower(name)
		i = strings.LastIndex(name, ".")
		if i == -1 {
			cache["unknown"] = true
		} else {
			cache[name[i:]] = true
		}
		name = ""
	}

	types = make([]string, 0, len(cache))
	for s := range cache {
		types = append(types, s)
	}
	return
}

// 目录拷贝
func DirCopy(src, dist string) {
	MkdirAll(dist)

	fileInfos, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, file := range fileInfos {
		fileSrc := src + FileSep + file.Name()
		fileDist := dist + FileSep + file.Name()
		if file.IsDir() {
			DirCopy(fileSrc, fileDist)
			continue
		}
		_, _ = FileCopy(fileSrc, fileDist)
	}
}

// 写入文本到指定文件
func WriteText(text, dist string) error {
	// 创建输出文件的父目录
	//MkdirParent(dist)
	//create, err := os.Create(dist)
	//if err != nil {
	//	_ = fmt.Errorf("写入文本到指定文件失败, err: %v", err)
	//	return
	//}
	//defer create.Close()
	//_, _ = create.WriteString(text)
	return WriteData([]byte(text), dist)
}

// 写入数据到指定文件
func WriteData(data []byte, dist string) error {
	// 创建输出文件的父目录
	MkdirParent(dist)
	create, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer create.Close()
	_, err = create.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 将一个文件读取成字符串返回
func ReadText(file string) (string, error) {
	bytes, err := ReadBytes(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 读取一个文件的 byte 二进制
func ReadBytes(file string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

// 清空一个目录的所有内容
func ClearDir(dir string) {
	_ = os.RemoveAll(dir)
}

// 从 from 到 to 的相对路径
//
// outJsPath := "/Users/ankang/git/saisheng/slgrpg/temp/quick-scripts/assets/script/feature/battleoverride"
// filePath := "/Users/ankang/git/saisheng/slgrpg/assets/script/feature/battleoverride"
// fmt.Println(outJsPath, filePath)
// path := fileutil.GetRelativePath(outJsPath, filePath)
// fmt.Println(path)
//
// Deprecated: use filepath.Rel(from, to) replace this method
func GetRelativePath(from, to string) string {
	var fromArr = strings.Split(from, FileSep)
	var toArr = strings.Split(to, FileSep)
	maxLen := len(fromArr)
	if toLen := len(toArr); toLen > maxLen {
		maxLen = toLen
	}
	sameLen := 0

	for i := 0; i < maxLen; i++ {
		if fromArr[i] != toArr[i] {
			break
		}
		sameLen++
	}

	sb := bytes.Buffer{}
	for i := 0; i < len(fromArr)-sameLen; i++ {
		sb.WriteString("../")
	}

	for i := 0; i < len(toArr)-sameLen; i++ {
		sb.WriteString(toArr[i+sameLen])
		sb.WriteString("/")
	}
	path := sb.String()
	return path[:len(path)-1]
}

// 获取md5
// Deprecated: Use commutil.GetMd5 instead.
func GetMd5(data *[]byte) string {
	return commutil.GetMd5(data)
}

// 生成base64
// Deprecated: Use commutil.EncodeBase64 instead.
func EncodeBase64(data *[]byte) string {
	return commutil.EncodeBase64(data)
}

// 解析base64
// Deprecated: Use commutil.DecodeBase64 instead.
func DecodeBase64(text string) ([]byte, error) {
	return commutil.DecodeBase64(text)
}
