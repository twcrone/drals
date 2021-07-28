package main

import (
	"bufio"
	"fmt"
	"os"
)

type alias struct {
	Key     string
	Command string
}

func main() {
	file, _ := os.Open("/Users/twcrone/.drals")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, eachLine := range text {
		a := parse(eachLine)
		fmt.Println(a)
	}
}

func parse(line string) alias {
	alias := alias{Key: "key", Command: "command"}

	return alias
}

func list() {

}

func add(directory string) {

}
