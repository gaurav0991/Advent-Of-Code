package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Reading file")

	// dat, err := os.ReadFile("input.txt")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	finalSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numberMul := 10
		secondDig := 0
		firstDig := 0
		// reading line by line digits
		for _, char := range line {
			// fmt.Println(char)
			if unicode.IsDigit(char) {
				if numberMul == 10 {
					numberMul = 1
					firstDig = int(char - '0')

				}
				secondDig = int(char - '0')
			}
		}
		//calculating final sum
		finalSum += ((firstDig * 10) + secondDig)
	}
	fmt.Printf("Sum of Part_1: %d\n", finalSum)

}
