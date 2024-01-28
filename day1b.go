package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// array of strings declare
var numberMap = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	fmt.Println("Program strating")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening file")
		return
	}
	defer file.Close()
	finakSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		foundFirst := false
		firstDig := 0
		lastDig := 0

		sum := 0

		for index, char := range line {
			if !foundFirst {

				for i, v := range numberMap {
					if strings.Contains(line[:index], v) {
						foundFirst = true
						firstDig = i
						// line = strings.ReplaceAll(line, v, s)

					}
				}
			}
			if foundFirst {
				for i, v := range numberMap {
					if strings.Contains(line[index:], v) {
						foundFirst = true
						lastDig = i
						// line = strings.ReplaceAll(line, v, s)

					}
				}
			}
			if unicode.IsDigit(char) {
				if !foundFirst {
					firstDig = int(char - '0')
					foundFirst = true
				}
				lastDig = int(char - '0')
			}
			sum = firstDig*10 + lastDig

		}
		finakSum += sum
		fmt.Println(finakSum)

	}
}
