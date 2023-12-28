package main

import (
	"fmt"
	"log"

	"github.com/Kirosich/OzonRoute256-2021/Lection_Modules/Storage/pkg/storage"
)

func main() {

	st := storage.NewStorage()
	file, err := st.Upload("test.txt", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client", file)
}
