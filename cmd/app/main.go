package main

import (
	"huff/internal/handler"
	"huff/internal/log"
)

func main() {
	log.InitLog()
	handler.Execute()
}
