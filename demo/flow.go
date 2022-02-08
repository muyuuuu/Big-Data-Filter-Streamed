package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Read(downstream chan string) {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		downstream <- fileScanner.Text()
	}
	close(downstream)
}

func filter(upstream, downstream chan string) {

	// item, ok := <-upstream
	// !ok means close
	for item := range upstream {
		tmp, err := strconv.ParseFloat(item, 32)
		if err == nil && tmp >= 0.5 {
			downstream <- item
		}
	}
	close(downstream)
}

func Finish(upstream chan string, write *bufio.Writer) {
	for item := range upstream {
		write.WriteString(item + "\n")
	}
	write.Flush()
}

func main() {
	start := time.Now()
	c0 := make(chan string)
	c1 := make(chan string)
	go Read(c0)
	go filter(c0, c1)
	// wait others goroutine
	Path := "result.txt"
	file, err := os.OpenFile(Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	Finish(c1, write)
	elapsed := time.Since(start)
	fmt.Println("cast, ", elapsed)
}
