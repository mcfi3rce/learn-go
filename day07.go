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
	IsFullHouse bool
	IsTwoPair   bool
	JokerCount  int
}

type KeyValue struct {
	Key   rune
	Value int
}

var cardValues = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14,
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
	var pairs []KeyValue
	for key, count := range cardCounts {
		pairs = append(pairs, KeyValue{Key: key, Value: count})
		if count > maxCount {
			maxCount = count
			keyType = key
		}
	}
	hand.CardCount = maxCount
	hand.HandType = keyType

	if hand.CardCount == 5 {
		return hand
	}
	// Figure out if it's a full house
	if maxCount == 3 {
		// Delete the key that is the three of a kind and see if the other two match
		remaining := strings.ReplaceAll(hand.Key, string(hand.HandType), "")
		if remaining[0] == remaining[1] {
			hand.IsFullHouse = true
		}
		return hand
	}

	twoPairs := 0
	for _, r := range cardCounts {
		if r == 2 {
			twoPairs++
		}
	}
	if twoPairs == 2 {
		hand.IsTwoPair = true
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	secondHighestCard := pairs[1].Key
	secondHighestCardCount := pairs[1].Value

	if strings.Contains(hand.Key, "J") {
		fmt.Println("Joker in hand", hand)
		jokerCount := cardCounts['J']
		fmt.Println("Joker Count", jokerCount)
		hand.JokerCount = jokerCount

		if hand.HandType == 'J' {
			fmt.Println("Joker is the highest card count", hand)
			hand.HandType = secondHighestCard
			hand.CardCount = secondHighestCardCount
			fmt.Println("Now the hand is", string(hand.HandType))
		}

		if hand.CardCount == 3 {
			if jokerCount == 1 {
				hand.IsFullHouse = true
				fmt.Println("Full House with Joker", hand)
			} else {
				hand.CardCount += jokerCount
				fmt.Println("Not a full house", hand)
			}
			return hand
		}
		if hand.CardCount == 2 {
			if jokerCount == 1 {
				if hand.IsTwoPair {
					hand.IsFullHouse = true
					fmt.Println("Was a Two Pair now a Full House with Joker", hand)
				} else {
					hand.IsTwoPair = true
					fmt.Println("Two Pair with Joker", hand)
				}
			} else {
				hand.CardCount += jokerCount
				fmt.Println("Not a two pair", hand)
			}
			return hand
		} else {
			hand.CardCount += jokerCount
			fmt.Println("Not a full house or two pair", hand)
			if hand.CardCount > 5 {
				fmt.Println("Hand has more than 5 cards", hand)
			}
			return hand
		}
	}
	fmt.Println("No hit end state", hand)
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

			if hands[i].CardCount == 2 {
				// FullHouse has higher priority
				if hands[i].IsTwoPair && !hands[j].IsTwoPair {
					return false
				}
				if !hands[i].IsTwoPair && hands[j].IsTwoPair {
					return true
				}
			}

			for k := 0; k < len(hands[i].Key); k++ {
				if cardValues[rune(hands[i].Key[k])] > cardValues[rune(hands[j].Key[k])] {
					//fmt.Println("Hand a: ", string(hands[i].Key[k]), " > Hand b: ", string(hands[j].Key[k]))
					//fmt.Println("Card Number: ", k+1, "Card Value: ", cardValues[rune(hands[i].Key[k])], "Card Value: ", cardValues[rune(hands[j].Key[k])])
					//fmt.Println("Hand:", hands[i], "Hand:", hands[j])
					return false
				}
				if cardValues[rune(hands[i].Key[k])] < cardValues[rune(hands[j].Key[k])] {
					//fmt.Println("Hand a: ", string(hands[i].Key[k]), " < Hand b: ", string(hands[j].Key[k]))
					//fmt.Println("Card Number: ", k+1, "Card Value: ", cardValues[rune(hands[i].Key[k])], "Card Value: ", cardValues[rune(hands[j].Key[k])])
					//fmt.Println("Hand:", hands[i], "Hand:", hands[j])
					return true
				}
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

func day07pt2(test []PokerHand) int {
	file, err := os.ReadFile("./day_07.txt")

	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")

	var hands []PokerHand
	if test == nil {

		// Start with setting the key to the hand
		// Then sort it in order of card counts then card values

		for _, line := range text {
			keyValue := strings.Fields(line)
			hand := extractHandValues(keyValue[0])
			hand.Value, _ = strconv.Atoi(keyValue[1])
			hands = append(hands, hand)
		}
	} else {
		hands = test
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

			if hands[i].CardCount == 2 {
				// FullHouse has higher priority
				if hands[i].IsTwoPair && !hands[j].IsTwoPair {
					return false
				}
				if !hands[i].IsTwoPair && hands[j].IsTwoPair {
					return true
				}
			}

			for k := 0; k < len(hands[i].Key); k++ {
				if cardValues[rune(hands[i].Key[k])] > cardValues[rune(hands[j].Key[k])] {
					//fmt.Println("Hand a: ", string(hands[i].Key[k]), " > Hand b: ", string(hands[j].Key[k]))
					//fmt.Println("Card Number: ", k+1, "Card Value: ", cardValues[rune(hands[i].Key[k])], "Card Value: ", cardValues[rune(hands[j].Key[k])])
					//fmt.Println("Hand:", hands[i], "Hand:", hands[j])
					return false
				}
				if cardValues[rune(hands[i].Key[k])] < cardValues[rune(hands[j].Key[k])] {
					//fmt.Println("Hand a: ", string(hands[i].Key[k]), " < Hand b: ", string(hands[j].Key[k]))
					//fmt.Println("Card Number: ", k+1, "Card Value: ", cardValues[rune(hands[i].Key[k])], "Card Value: ", cardValues[rune(hands[j].Key[k])])
					//fmt.Println("Hand:", hands[i], "Hand:", hands[j])
					return true
				}
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
	return total
}
