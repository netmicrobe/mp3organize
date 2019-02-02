package main

import (
  "fmt"
  "log"

  "github.com/bogem/id3v2"
)

func main() {
  // Open file and parse tag in it.
  tag, err := id3v2.Open("Out of My Hands.mp3", id3v2.Options{Parse: true})
  if err != nil {
     log.Fatal("Error while opening mp3 file: ", err)
  }
  defer tag.Close()

  // Read frames.
  fmt.Println(tag.Artist())
  fmt.Println(tag.Title())
}