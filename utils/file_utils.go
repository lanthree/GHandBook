package utils

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// File 文件
type File struct {
	Name           string
	Content        string
	Path           string
	LastModifyTime int64
}

// Dir 目录
type Dir struct {
	Name           string
	Path           string
	LastModifyTime int64
}

// IsFileExit 判断文件（非目录）是否存在
func IsFileExit(path string) (bool, error) {

	stat, err := os.Stat(path)
	if err == nil {
		if stat.IsDir() {
			return false, nil
		}
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	// 其他失败
	return true, err
}

// GetFileModTime 获取文件/目录修改事件
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

// GetFilesAndDirsNotRecursive 获取指定目录下文件和目录 不递归
func GetFilesAndDirsNotRecursive(path string) (files []File, dirs []Dir, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	pathSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, Dir{fi.Name(), path + pathSep + fi.Name(), fi.ModTime().Unix()})
		} else {
			file := File{fi.Name(), "", path + pathSep + fi.Name(), fi.ModTime().Unix()}

			f, err := os.Open(file.Path)
			if err != nil {
				log.Printf("open file:%s fail:%v\n", file.Path, err)
				continue
			}
			defer f.Close()

			bytes, err := ioutil.ReadAll(f)
			file.Content = string(bytes)
			if err != nil {
				log.Printf("read file:%s fail:%v\n", file.Path, err)
				continue
			}
			files = append(files, file)
		}
	}

	return files, dirs, nil
}
