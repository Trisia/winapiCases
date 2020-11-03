package main

import (
	"crypto/rand"
	"os"
)

func main() {
	file, err := os.OpenFile("BigFile.dat", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 1 KB
	// 1024 KB = 1 MB
	// 1024 * 1024 KB = 1 GB
	size := 1024 * 1024
	for i := 0; i < size; i++ {
		// 1KB
		buff := make([]byte, 1024)
		rand.Reader.Read(buff)
		file.Write(buff)
	}
}
