package service

import (
	"fmt"
	"texasPoker/model"
	"testing"
)

func TestCardIsStraight(t *testing.T){
	if isOk,err := CardIsStraight("23456"); err != nil {
		fmt.Printf("非顺子\n")
	} else {
		if isOk {
			fmt.Printf("顺子\n")
		}
	}
}

func TestFiveCardParse(t *testing.T) {
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src:"6cAc6h4hQd"})
	PokerList = append(PokerList, model.HandCards{Src:"9h9c5cJcTh"})

	for _, alice := range PokerList {
		FiveCardParse(&alice)
		fmt.Println(alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestCompareTwoHandCard2(t *testing.T) {
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src:"2s6cKcJh2h"})
	PokerList = append(PokerList, model.HandCards{Src:"JcXn7s4d3s"})

	CompareTwoHandCard(&PokerList[0], &PokerList[1])

	fmt.Println(PokerList[0], " handcard type: ", model.HandCardType[PokerList[0].Type])
	fmt.Println(PokerList[1], " handcard type: ", model.HandCardType[PokerList[1].Type])

}


