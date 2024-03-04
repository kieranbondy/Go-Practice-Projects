package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(strings.Join(os.Args[1:], "")) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
