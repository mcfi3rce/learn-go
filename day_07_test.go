package main

import (
	"testing"
)

func TestExtractHandValuesWithFullHouse(t *testing.T) {
	hand := extractHandValues("3344J")
	if !hand.IsFullHouse {
		t.Errorf("Expected hand to be a FullHouse, but it was not")
	}
}

func TestExtractHandValuesWithTwoPair(t *testing.T) {
	hand := extractHandValues("3434A")
	if !hand.IsTwoPair {
		t.Errorf("Expected hand to be a TwoPair, but it was not")
	}
}

func TestExtractHandValuesWithJokerFullHouse(t *testing.T) {
	hand := extractHandValues("AA3J3")
	if !hand.IsFullHouse || hand.JokerCount != 1 || hand.CardCount != 2 {
		t.Errorf("Expected hand to be a FullHouse with 1 Jokers, but got %v with %d Jokers", hand.IsFullHouse, hand.JokerCount)
	}
}

func TestExtractHandValuesWithHighCard(t *testing.T) {
	hand := extractHandValues("A4562")
	if hand.CardCount != 1 {
		t.Errorf("Expected highest card count to be 1, but got %d", hand.CardCount)
	}
}

func TestDay07pt2SortingLogic(t *testing.T) {
	hands := []PokerHand{
		{Key: "32T3K", Value: 765},
		{Key: "T55J5", Value: 684},
		{Key: "KK677", Value: 28},
		{Key: "KTJJT", Value: 220},
		{Key: "QQQJA", Value: 483},
	}
	var value = day07pt2(hands) // Assuming this function sorts `hands` slice
	if value != 5905 {
		t.Errorf("Sorting logic failed to correctly sort hands based on the given criteria")
	}
}

func TestExtractHandValuesWithInvalidHand(t *testing.T) {
	hand := extractHandValues("ZZZZZ")
	if hand.Key != "ZZZZZ" {
		t.Errorf("Expected hand key to be 'ZZZZZ', but got %s", hand.Key)
	}
}
