package main

import (
	"go.nndi.cloud/loyo"
)

func main() {
	if err := loyo.NewServer(":8000"); err != nil {
		panic(err)
	}
}
