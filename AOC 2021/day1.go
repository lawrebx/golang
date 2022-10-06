//https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var prevDepth = 0

	for scanner.Scan() {
		var curDepth = scanner.Text()

	}
}
