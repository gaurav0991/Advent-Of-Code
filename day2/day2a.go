package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Program starting")
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	gamerCounter := 1
	possibleYes := 0
	for scanner.Scan() {
		line := scanner.Text()
		arrDiv := strings.Split(line, ":")
		possiblities := strings.Split(arrDiv[1], ";")
		possibleNo := false
		for _, v := range possiblities {
			// fmt.Println(v)
			// splitting each comma to get the array
			eachCube := strings.Split(v, ",")
			for _, cube := range eachCube {
				result := strings.ReplaceAll(cube, " ", "")
				// fmt.Println(result)
				numberPres := 0
				for _, ch := range result {
					if unicode.IsDigit(ch) {
						if numberPres != 0 {
							numberPres *= 10
						}
						numberPres += int(ch - '0')
					}
				}

				if strings.Contains(result, "blue") {
					// fmt.Printf("blue %d\n", numberPres)
					if numberPres > 14 {
						possibleNo = true
					}
				}
				if strings.Contains(result, "red") {
					// fmt.Printf("red %d\n", numberPres)

					if numberPres > 12 {
						possibleNo = true
					}
				}
				if strings.Contains(result, "green") {
					// fmt.Printf("green %d\n", numberPres)

					if numberPres > 13 {
						possibleNo = true
					}
				}
			}
		}
		if !possibleNo {
			fmt.Println(gamerCounter)
			possibleYes += gamerCounter
		}
		gamerCounter += 1
	}
	fmt.Println(possibleYes)

}
