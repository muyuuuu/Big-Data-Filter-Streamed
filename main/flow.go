package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func parse_string_bound(item string, col int, limit relation, string_bound string) bool {
	tmp := strings.Split(item, ",")
	if !strings.Contains(tmp[col], string_bound) {
		return true
	}
	return false
}

func parse_number_bound(item string, col int, limit relation, number_bound string) bool {

	tmp := strings.Split(item, ",")
	// fmt.Println(tmp[col], number_bound)
	// os.Exit(0)
	switch limit {
	case equal:
		if tmp[col] == number_bound {
			return true
		}
	case large:
		if tmp[col] >= number_bound {
			return true
		}
	case small:
		if tmp[col] < number_bound {
			return true
		}
	}
	return false
}

func filter(upstream, downstream chan string, restriction []limitation) {
	var length = len(restriction)
	if length != 0 {
		for item := range upstream {
			var flag = true
			for _, val := range restriction {
				col := val.col
				limit := val.limit
				string_bound := val.string_bound
				number_bound := val.number_bound
				if string_bound != "none" {
					if !parse_string_bound(item, col, limit, string_bound) {
						flag = false
					}
				}
				if number_bound != "none" {
					if !parse_number_bound(item, col, limit, number_bound) {
						flag = false
					}
				}
			}
			if flag == true {
				downstream <- item
			}
		}
	} else {
		for item := range upstream {
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

type relation string

const equal relation = "equal"
const large relation = "large"
const small relation = "small"
const contain relation = "contain"

type limitation struct {
	col          int
	limit        relation
	string_bound string
	number_bound string
}

func ParseLimitation() []limitation {
	var all []limitation

	// 性别为 1
	tmp := limitation{
		col:          1,
		limit:        equal,
		string_bound: "none",
		number_bound: "1",
	}
	all = append(all, tmp)
	// 考核分数大于 90
	tmp = limitation{
		col:          2,
		limit:        large,
		string_bound: "none",
		number_bound: "90",
	}
	all = append(all, tmp)
	return all
}

func main() {
	start := time.Now()
	c0 := make(chan string)
	c1 := make(chan string)
	go Read(c0)
	go filter(c0, c1, ParseLimitation())

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
