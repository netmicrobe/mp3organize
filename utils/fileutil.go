package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type func_type_opmp3 func(string)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil{
		return false, err
	}
	return fileInfo.IsDir(), err
}


func IsMp3File(path string) (bool, error) {
	r, err := regexp.MatchString(".mp3", path)
	if err != nil{
		return false, err
	}
	return r, err
}


func RenameFile(path string, newname string) {
	reg, err := regexp.Compile("[\\/\\*:\\|\"<> &\\(\\)]+")
	if err != nil {
		log.Fatal(err)
	}
	newname = reg.ReplaceAllString(newname, "_")

	newpath := filepath.Dir(path) + string(os.PathSeparator) + newname
	fmt.Println("newpath = " , newpath)
	
	err = os.Rename(path, newpath)
	if err != nil {
		log.Fatal(err)
	}
}

// path 参数可以是文件夹或文件
func EachFiles(path string, ope func_type_opmp3) {
	// 判断是 文件 还是 文件夹
	isDir, err := IsDirectory(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if isDir {
		// do directory stuff
		fmt.Println("directory : ", path)
		err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					fmt.Println(path, info.Size())
					r, err := IsMp3File(info.Name())
					if err == nil && r {
						ope(path)
					}
				}
				return nil
			})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// do file stuff
		fmt.Println("file : ", path)
		ope(path)
	}
}







