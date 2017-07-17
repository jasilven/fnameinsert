package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("%v inserts file name (excluding extension) to the beginning of the file\n", os.Args[0])
		fmt.Println("Usage:", os.Args[0], "file1 file2...")
		os.Exit(0)
	}

	for _, fname := range os.Args[1:] {
		fmt.Printf("Processing %v: ", fname)

		data, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalln(err.Error())
		}

		name := strings.Split(fname, ".")[0]
		b := []byte(name + "\n")

		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		firstline := ""
		if scanner.Scan() {
			firstline = scanner.Text()
		}
		if strings.Compare(firstline, name) == 0 {
			fmt.Println("Already OK")
			continue
		}

		err = ioutil.WriteFile(fname, append(b[:], data[:]...), 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println("Done")
	}
}
