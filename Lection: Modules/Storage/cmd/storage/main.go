package main

import (
	"fmt"
	"log"

	"github.com/Kirosich/OzonRoute256-2021/LectionModules/Storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
