package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PokerHand struct {
	Key         string
	Value       int
	CardCount   int
	HandType    rune
	HighCard    int
	IsFullHouse bool
}

var cardValues = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func extractHandValues(key string) PokerHand {
	var hand PokerHand
	hand.Key = key

	// Count the occurrences of each card
	cardCounts := make(map[rune]int)
	for _, card := range hand.Key {
		cardCounts[card]++
	}

	// Find the maximum count value
	maxCount := 0
	keyType := rune(0)
	for key, count := range cardCounts {
		if count > maxCount {
			maxCount = count
			keyType = key
		}
	}
	hand.CardCount = maxCount
	hand.HandType = keyType

	// Figure out if it's a full house
	if maxCount == 3 {
		// Delete the key that is the three of a kind and see if the other two match
		remaining := strings.ReplaceAll(hand.Key, string(hand.HandType), "")
		if remaining[0] == remaining[1] {
			hand.IsFullHouse = true
		}
	}

	// Find the high card
	var remaining = strings.ReplaceAll(hand.Key, string(hand.HandType), "")
	for i := 0; i < len(remaining); i++ {
		if cardValues[rune(remaining[i])] > hand.HighCard {
			hand.HighCard = cardValues[rune(remaining[i])]
		}
	}

	return hand
}

func day07pt1() {
	file, err := os.ReadFile("./day_07.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")

	var hands []PokerHand

	// Start with setting the key to the hand
	// Then sort it in order of card counts then card values

	for _, line := range text {
		keyValue := strings.Fields(line)
		hand := extractHandValues(keyValue[0])
		hand.Value, _ = strconv.Atoi(keyValue[1])
		hands = append(hands, hand)
	}

	// Sort the hands by card count then high card
	// Assuming hands is a slice of PokerHand
	// Note this is sorted lowest to highest so if it's higher rank return false
	sort.Slice(hands, func(i, j int) bool {
		// If neither or both are FullHouse, sort by CardCount, then by HighCard
		if hands[i].CardCount == hands[j].CardCount {

			if hands[i].CardCount == 3 {
				// FullHouse has higher priority
				if hands[i].IsFullHouse && !hands[j].IsFullHouse {
					return false
				}
				if !hands[i].IsFullHouse && hands[j].IsFullHouse {
					return true
				}
			}

			// Return the higher of the two card counts
			if cardValues[hands[i].HandType] > cardValues[hands[j].HandType] {
				fmt.Println("Hand:", hands[i], "Higher than Hand:", hands[j])
				fmt.Println("Card Values:", cardValues[hands[i].HandType], "VS", cardValues[hands[j].HandType])
				return false
			}

			if cardValues[hands[i].HandType] < cardValues[hands[j].HandType] {
				fmt.Println("Hand:", hands[j], "Lower than Hand:", hands[i])
				fmt.Println("Card Values:", cardValues[hands[i].HandType], "VS", cardValues[hands[j].HandType])
				return true
			}

			fmt.Println("USING HIGH CARD", hands[i], "VS", hands[j])
			if hands[i].HighCard < hands[j].HighCard {
				fmt.Println("Hand:", hands[i], "Lower than Hand:", hands[j])
				fmt.Println("High Cards:", hands[i].HighCard, "VS", hands[j].HighCard)
				return true
			}

			if hands[i].HighCard > hands[j].HighCard {
				fmt.Println("Hand:", hands[i], "Higher than Hand:", hands[j])
				fmt.Println("High Cards:", hands[i].HighCard, "VS", hands[j].HighCard)
				return false
			}

		}
		return hands[i].CardCount < hands[j].CardCount || hands[j].CardCount > hands[i].CardCount
	})

	total := 0
	for i, hand := range hands {
		total += hand.Value * (i + 1)
		fmt.Println("Hand:", hand, "Value:", hand.Value*(i+1), "Rank: ", i+1, "Total:", total)
	}

	fmt.Println("Total:", total)
}

func day07pt2() {
	file, err := os.ReadFile("./day_07.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")
	fmt.Println(text)

}
