package main

import (
	"fmt"
	"github.com/asorokovikov/storage/internal/storage"
	"log"
)

func main() {
	store := storage.NewStorage()

	file, err := store.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file uploaded", file)

	file, err = store.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file received", file)
}
