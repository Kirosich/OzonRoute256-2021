package main

import (
	"fmt"
	"log"

	"github.com/Kirosich/OzonRoute256-2021/Lection_Modules/Storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file, restoredFile)
}
