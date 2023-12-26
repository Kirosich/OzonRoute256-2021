package main

import (
	"fmt"
	"ozonmodules/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println(st)
}
