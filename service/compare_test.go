package service

import (
	"fmt"
	"texasPoker/model"
	"testing"
)

func TestCompareStraightFlush(t *testing.T) {
	card1 := "TJQKA"
	card2 := "9TJQK"
	result := CompareStraightFlush(card1, card2)

	fmt.Printf("result: %d\n",result)
}

func TestCompareTwoHandCard(t *testing.T) {
	alice := model.HandCards{Src: "QdQd8c4s8c"}
	bob := model.HandCards{Src: "Qc9sAcAsQh"}

	result := CompareTwoHandCard(&alice, &bob)

	fmt.Printf("alice[%s] - bob[%s], result: %d\n", alice.Src, bob.Src, result)
}

func TestCompareOnePair(t *testing.T) {
	a := "226JK"
	b := "347JX"
	result := CompareOnePair(a, b)

	fmt.Println(result)
}

func TestCompareTwoPair(t *testing.T) {
	a := "66JQQ"
	b := "66JJQ"
	result := CompareTwoPair(a, b)

	fmt.Println(result)
}