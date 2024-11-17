package main

import "huff/internal/log"

func main() {
	log := log.InitLog()
	log.Info("Text")
}
