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

func checkAroundForSymbol(r int, c int, text []string, symbols []string) (bool, []string) {
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
				symbols = append(symbols, val)
				return true, symbols
			}
		} else {
			//fmt.Println("Skipping out-of-bounds access for", newR, newC)
		}
	}
	return false, symbols
}

func main() {
	file, err := os.ReadFile("./day_03.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")
	total := 0
	var s []int
	var symbols []string

	for r := 0; r < len(text); r++ {
		num := ""
		isValid := false
		for c := 0; c < len(text[r]); c++ {
			for (text[r][c] != '.' && !isSymbol(string(text[r][c]))) && c < len(text[r])-1 {
				val := string(text[r][c])

				num += val
				// See if current number is a number
				if _, err := strconv.Atoi(val); err == nil {
					// Check all directions in the array to see if any of them are symbols
					// we only want to check if it's false otherwise we can just keep moving forward
					if !isValid {
						// if it is a symbol that means the number is valid
						isValid, symbols = checkAroundForSymbol(r, c, text, symbols)
					}
				}

				c++
			}
			if isValid {
				if x, err := strconv.Atoi(num); err == nil {
					total += x
					s = append(s, x)
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

	for i := 0; i < len(s) && i < len(symbols); i++ {
		fmt.Println(s[i], symbols[i])
	}
	fmt.Println("Total: ", total)
}
