package service

import (
	"fmt"
	"myproject/texasPoker/model"
	"testing"
)

func TestSevenCardHasStraight(t *testing.T) {
	sevenCardFace := "2333456"
	isOk := SevenCardHasStraight(sevenCardFace)
	if isOk {
		fmt.Printf("[%s] 是顺子", sevenCardFace)
	} else {
		fmt.Printf("[%s] 不是顺子", sevenCardFace)
	}
}

func TestSevenCardParse(t *testing.T) {
	alice := model.HandCards{
		SortFace: 	"A556789",
		IsGhost:	false,
	}
	SevenCardParse(&alice)
	fmt.Println(model.HandCardType[alice.Type])
}

func TestSevenCardHasFlush(t *testing.T) {
	alice := model.HandCards{

	}
	SevenCardHasFlush(&alice)
}
