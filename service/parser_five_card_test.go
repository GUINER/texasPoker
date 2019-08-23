package service

import (
	"fmt"
	"testing"
	"texasPoker/model"
)

func TestCardIsStraight(t *testing.T) {
	if isOk, err := CardIsStraight("23456"); err != nil {
		fmt.Printf("非顺子\n")
	} else {
		if isOk {
			fmt.Printf("顺子\n")
		}
	}
}

func TestFiveCardParse(t *testing.T) {
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "6cAc6h4hQd"})
	PokerList = append(PokerList, model.HandCards{Src: "9h9c5cJcTh"})

	for _, alice := range PokerList {
		FiveCardParse(&alice)
		fmt.Println(alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}
