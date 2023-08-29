package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Data []float64

const (
	empty = 0
)

func ParseData() {
	if len(os.Args) == 2 {
		data, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		in := bufio.NewScanner(data)
		for in.Scan() {
			var data float64
			fmt.Sscan(in.Text(), &data)
			Data = append(Data, data)
		}
		if len(Data) == empty {
			log.Fatal("Empty file")
		}
	} else {
		log.Fatal("Incorrect input:USE go run . file.txt")
	}
}
