package utils

import (
	"fmt"

	"github.com/bogem/id3v2"
)


func GetMp3NewFilename(mp3fi string) string {
	tag, err := id3v2.Open(mp3fi, id3v2.Options{Parse: true})
	if err != nil {
		fmt.Println("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read frames.
	newname := tag.Title() + "-" + tag.Artist() + ".mp3"
	
	return newname
}

func Renamemp3(mp3fi string) {
	newname := GetMp3NewFilename(mp3fi)
	RenameFile(mp3fi, newname)
}

func PrintInfoMp3(mp3fi string) {
	tag, err := id3v2.Open(mp3fi, id3v2.Options{Parse: true})
	if err != nil {
		fmt.Println("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	fmt.Println("== ", mp3fi)
	fmt.Println("   SONG    : " , tag.Title())
	fmt.Println("   ARTIST  : " , tag.Artist())
	fmt.Println()
}






