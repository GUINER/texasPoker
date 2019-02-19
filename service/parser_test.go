package service

import (
	"fmt"
	"myproject/texasPoker/model"
	"testing"
)

//=========================7张牌测试=======================
//是否有顺子
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
		//SortFace: 	"TJQKKAA",
		IsGhost:	false,
		Src: "AdKcJcQcTcKsAs",
	}
	SortCard(&alice)
	SevenCardParse(&alice)
	fmt.Println(alice)
	fmt.Println(model.HandCardType[alice.Type])
}

// 测试7张牌是否有同花
func TestSevenCardHasFlush(t *testing.T) {
	alice := model.HandCards{
		Src:"XnAc5cTcKh6d9h",
		IsGhost:true,
	}
	SortCard(&alice)
	if SevenCardHasFlush(&alice) {
		alice.Type = model.FlUSH
	}
	fmt.Println(alice)
	fmt.Println(model.HandCardType[alice.Type])
}


func TestSevenCardFourOfAKind(t *testing.T) {
	alice := &model.HandCards{
		Src: "AcAdAs7h7dXn7c",
		IsGhost: true,
	}
	SortCard(alice)
	if SevenCardFourOfAKind(alice) {
		alice.Type = model.FOUROFAKIND
	}
	fmt.Println(*alice)
	fmt.Println(model.HandCardType[alice.Type])
}


//=========================7张牌测试end=======================
