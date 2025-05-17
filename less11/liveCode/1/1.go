package main

import (
	"fmt"
)

type MediaPlayer interface {
	Play(filename string)
}

type AudioPlayer struct {
	fileName string
}

func (t AudioPlayer) Play(some string) {
	fmt.Println(some, t.fileName)
}

type VideoPlayer struct {
	fileName string
}

func (m VideoPlayer) Play(some string) {
	fmt.Println(some, m.fileName)
}

func play(m []MediaPlayer) {
	for _, v := range m {
		v.Play("someFile")
	}
}

func main() {

	play([]MediaPlayer{
		VideoPlayer{fileName: "m4a"},
		AudioPlayer{fileName: "mp3"},
		VideoPlayer{fileName: "mkv"},
		AudioPlayer{fileName: "tt"},
	})
}
