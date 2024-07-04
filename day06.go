package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day06pt1() {
	file, err := os.ReadFile("./day_06.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")
	time := strings.Fields(text[0])[1:]
	distance := strings.Fields(text[1])[1:]
	totalOptions := 1

	fmt.Println(time)
	fmt.Println(distance)
	for i := 0; i < len(time); i++ {
		totalTime, _ := strconv.Atoi(time[i])
		threshold, _ := strconv.Atoi(distance[i])
		count := 0
		for j := 1; j < totalTime; j++ {
			remainingTime := totalTime - j
			speed := j
			distance := speed * remainingTime
			if distance > threshold {
				fmt.Println("Time: ", totalTime, "Speed: ", speed, "RemainingTime: ", remainingTime, "Distance: ", distance)
				count++
			}
		}
		totalOptions *= count
	}
	fmt.Println(totalOptions)
}

func day06pt2() {
	file, err := os.ReadFile("./day_06.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")
	timeStr := strings.Join(strings.Fields(text[0])[1:], "")
	distanceStr := strings.Join(strings.Fields(text[1])[1:], "")

	time, _ := strconv.Atoi(timeStr)
	threshold, _ := strconv.Atoi(distanceStr)

	count := 0
	for i := 1; i < time; i++ {
		remainingTime := time - i
		speed := i
		distance := speed * remainingTime
		if distance < threshold {
			count++
		} else {
			break
		}
	}
	for j := time; j > 0; j-- {
		remainingTime := time - j
		speed := j
		distance := speed * remainingTime
		if distance < threshold {
			count++
		} else {
			break
		}
	}

	fmt.Println(time - count)

}
