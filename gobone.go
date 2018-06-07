package gobone

import (
	"fmt"
	"log"
	"os"
)

func WriteAndClose(fi string, v string) {
	f, err := os.OpenFile(fi, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(v)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write to %s\n", fi)
}

func ReadByteAndClose(fi string) []byte {

	var size int
	v := make([]byte, 1)
	f, err := os.OpenFile(fi, os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	size, err = f.Read(v)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %v from %s\nWas %d bytes in len", v, fi, size)
	return v
}
