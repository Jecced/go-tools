package fileutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// FileSep 系统路径分隔符
	FileSep = string(os.PathSeparator)
)

// MkdirAll 根据路径创建文件夹
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

// MkdirParent 创建一个文件的父目录
func MkdirParent(path string) error {
	parent := GetParentDir(path)
	if !PathExists(parent) {
		return MkdirAll(parent)
	}
	return nil
}

// FindBySuffix 获取某个文件夹下所有指定后缀的文件
func FindBySuffix(dirPath string, suffix string) (files []string, err error) {
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
			newFiles, _ := FindBySuffix(dirPath+FileSep+fi.Name(), suffix)
			files = append(files, newFiles...)
		} else if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPath+FileSep+fi.Name())
		}
	}
	return files, nil
}

// Deprecated: use FindBySuffix(from, to) replace this method
func GetFilesBySuffix(dirPath string, suffix string) (files []string, err error) {
	return FindBySuffix(dirPath, suffix)
}

// FileCopy 文件拷贝
func FileCopy(src, dist string) (err error) {
	_ = os.Remove(dist)

	// 开启 源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		_ = srcFile.Close()
	}()

	// 创建输出文件的父目录
	err = MkdirParent(dist)
	if err != nil {
		return err
	}

	// 创建目标文件
	distFile, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer func() {
		_ = distFile.Close()
	}()

	// 创建缓冲区
	bs := make([]byte, 1024*10, 1024*10)
	n := -1
	for {
		n, err = srcFile.Read(bs)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return err
		}
		_, _ = distFile.Write(bs[:n])
	}
	return nil
}

// DirCopy 目录拷贝
func DirCopy(src, dist string) error {
	err := MkdirAll(dist)
	if err != nil {
		return err
	}

	fileInfos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range fileInfos {
		fileSrc := src + FileSep + file.Name()
		fileDist := dist + FileSep + file.Name()
		if file.IsDir() {
			DirCopy(fileSrc, fileDist)
			continue
		}
		_ = FileCopy(fileSrc, fileDist)
	}
	return nil
}

// PathExists 判断一个路径是否存在
func PathExists(path string) bool {
	stat, _ := os.Stat(path)
	return stat != nil
}

// GetParentDir 获取一个路径的父目录地址
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

// GetFileName 根据路径, 获取文件名
func GetFileName(path string) string {
	_, name := filepath.Split(path)
	return name
}

// GetSuffix 根据路径, 获取文件后缀
func GetSuffix(path string) string {
	name := GetFileName(path)
	lastIndex := strings.LastIndex(name, ".")
	if -1 == lastIndex {
		return "unknown"
	}
	return name[lastIndex:]
}

// PathFormat 路径格式化, 标准化一个路径到当前系统规范
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

// FindAllFileTypes 获取目录下所有文件类型
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
			out := FindAllFileTypes(dir + FileSep + info.Name())
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

// WriteText 写入文本到指定文件
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

// WriteData 写入数据到指定文件
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

// ReadText 将一个文件读取成字符串返回
func ReadText(file string) (string, error) {
	bytes, err := ReadBytes(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadBytes 读取一个文件的 byte 二进制
func ReadBytes(file string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

// ClearDir 清空一个目录的所有内容
func ClearDir(dir string) {
	_ = os.RemoveAll(dir)
}

// RemoveFile 可以删除一个文件, 空文件夹
func RemoveFile(file string) error {
	return os.Remove(file)
}

// GetRelativePath 从 from 到 to 的相对路径
//
// old := "/Users/xxx/yyy/zzz/www/temp/quick-scripts/assets/script/feature/battle"
// now := "/Users/xxx/yyy/zzz/www/assets/script/feature/battle"
// path := GetRelativePath(old, now)
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

// DelEmptyDir 自内向外删除所有空文件夹, 如果文件是.DS_Store的话, 也会一起删除
func DelEmptyDir(dir string) (bool, error) {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}

	// 确定文件夹内部文件数量
	l := len(list)
	path := ""

	for _, info := range list {
		path = dir + string(os.PathSeparator) + info.Name()
		if !info.IsDir() && info.Name() == ".DS_Store" {
			err := os.Remove(path)
			if err != nil {
				continue
			}
			l--
			continue
		}
		if !info.IsDir() {
			continue
		}
		b, err := DelEmptyDir(path)
		if err != nil {
			continue
		}
		if b {
			err := os.Remove(path)
			if err != nil {
				continue
			}
			l--
		}
	}
	return l == 0, nil
}

// FileSize 格式化文件的大小
func FileSize(path string) (size string) {
	info, err := os.Stat(path)
	if err != nil {
		return err.Error()
	}
	fileSize := info.Size()
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
