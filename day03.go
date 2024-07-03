package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSymbol(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' {
			return false // Contains a letter or digit, not a symbol
		}
	}
	return true // No letters or digits found, so it's a symbol string
}

func checkAroundForSymbol(r int, c int, text []string, num Number) (bool, Number) {
	/*
	  467-1,14..
	  ..0,-1*......
	  ..35.
	*/
	points := [][]int{{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}

	for _, dir := range points {
		newR := r + dir[0]
		newC := c + dir[1]
		if newR >= 0 && newC >= 0 && newR < len(text) && newC < len(text[newR]) {
			val := string(text[newR][newC])
			if isSymbol(val) {
				return true, Number{-1, val, newR, newC}
			}
		}
	}
	return false, Number{-1, "", -1, -1}
}

type Number struct {
	value     int
	symbol    string
	symbolRow int
	symbolCol int
}

func day03part1() {
	file, err := os.ReadFile("./day_03.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")
	total := 0

	var numbers []Number

	for r := 0; r < len(text); r++ {
		num := ""
		isValid := false
		number := Number{}
		for c := 0; c < len(text[r]); c++ {
			for c < len(text[r]) {
				val := string(text[r][c])

				if val == "." || isSymbol(val) {
					break
				}

				num += val
				// See if current number is a number
				if _, err := strconv.Atoi(val); err == nil {
					// Check all directions in the array to see if any of them are symbols
					// we only want to check if it's false otherwise we can just keep moving forward
					if !isValid {
						// if it is a symbol that means the number is valid
						isValid, number = checkAroundForSymbol(r, c, text, number)
					}
				}

				c++
			}
			if isValid {
				if x, err := strconv.Atoi(num); err == nil {
					total += x
					number.value = x
					numbers = append(numbers, number)
				}
			}
			isValid = false
			num = ""
		}
	}
	//fmt.Printf("%c", text[0][0])
	/**
	1. Iterate through the array.
	2. If you find a number check it's four directions for a symbol
		a. if you don't find a symbol and the next character is still a number add it
		b. if you haven't found any symbols once you hit the end of a number don't add it
	**/

	// iterate through the array and if two numbers have a * and the same row column values
	// then multiply them together
	// Map to keep track of numbers by their location
	locationMap := make(map[string][]int)

	// Populate the map
	for _, num := range numbers {
		if num.symbol == "*" {
			key := fmt.Sprintf("%d-%d", num.symbolRow, num.symbolCol)
			locationMap[key] = append(locationMap[key], num.value)
		}
	}

	productTotal := 0
	for key, vals := range locationMap {
		product := 1
		if len(vals) > 1 {
			for _, val := range vals {
				product *= val
			}
			fmt.Printf("Location %s has product: %d\n", key, product)
		}
		if product > 1 {
			productTotal += product
		}
	}

	fmt.Println(numbers)
	fmt.Println("Total: ", total)
	fmt.Println("Product Total: ", productTotal)
}
