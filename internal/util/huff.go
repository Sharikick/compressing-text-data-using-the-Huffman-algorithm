package util

import "fmt"

func Code(path string) {
	text := string(ReadFile(path))
	fmt.Println(text)
}

func Decode() {}
