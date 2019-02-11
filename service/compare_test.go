package service

import (
	"fmt"
	"myproject/texasPoker/model"
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
