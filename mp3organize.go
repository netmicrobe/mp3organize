package main

import (
  "fmt"
//  "log"
  "flag"
  "os"
  "path/filepath"
  "regexp"

  "github.com/bogem/id3v2"
)

func main() {
  recursivePtr := flag.Bool("r", false, "recursive in sub-folders")
  flag.Parse()
  target := flag.Arg(0) // 文件，或者文件夹
  if target == "" {
    fmt.Println("错误！需要指明文件 或 文件夹的路径！")
    return
  }
  
  fmt.Println(*recursivePtr)
  fmt.Println(target)

  // 判断是 文件 还是 文件夹
  isDir, err := IsDirectory(target)
  if err != nil {
    fmt.Println(err)
    return
  }
  if isDir {
    // do directory stuff
    fmt.Println("directory : ", target)
    err := filepath.Walk(target,
      func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
          fmt.Println(path, info.Size())
          r, err := regexp.MatchString(".mp3", info.Name())
          if err == nil && r {
            renamemp3(path)
          }
        }
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
  } else {
    // do file stuff
    fmt.Println("file : ", target)
    renamemp3(target)
  }

}

func IsDirectory(path string) (bool, error) {
  fileInfo, err := os.Stat(path)
  if err != nil{
    return false, err
  }
  return fileInfo.IsDir(), err
}

func renamemp3(mp3fi string) {
  tag, err := id3v2.Open(mp3fi, id3v2.Options{Parse: true})
  if err != nil {
     fmt.Println("Error while opening mp3 file: ", err)
  }
  defer tag.Close()

  // Read frames.
  fmt.Println("    TRY TO RENAME : ", tag.Artist(), "-", tag.Title())
}








