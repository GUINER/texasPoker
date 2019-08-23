package service

import (
	"fmt"
	"testing"
	"texasPoker/model"
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
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAs7h7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7d9h7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AdKcJcQcTcKsAs"})
	//alice := model.HandCards{
	//	//SortFace: 	"TJQKKAA",
	//	IsGhost:	false,
	//	Src: "AdKcJcQcTcKsAs",
	//}
	//SortCard(&alice)
	//SevenCardParse(&alice)
	//fmt.Println(alice)
	//fmt.Println(model.HandCardType[alice.Type])
	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		SevenCardParse(&alice)

		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

// 测试7张牌是否有同花
func TestSevenCardHasFlush(t *testing.T) {
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "XnAc5cTcKh6d9h"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAc5cTcKc6d9h"})
	PokerList = append(PokerList, model.HandCards{Src: "9cAc5cTcKc6d9h"})
	//PokerList = append(PokerList, model.HandCards{Src:"XnAc5cTcKh6d9h",IsGhost:true,})
	//PokerList = append(PokerList, model.HandCards{Src:"XnAc5cTcKh6d9h",IsGhost:true,})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardHasFlush(&alice) {
			alice.Type = model.FlUSH
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardFourOfAKind(t *testing.T) {
	//4+X,3+1+X,4+1
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAs7h7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7d9h7c"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardFourOfAKind(&alice) {
			alice.Type = model.FOUROFAKIND
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardFullHouseAndThreeKind(t *testing.T) {
	// 3+2,3+2+X or 3+1+1,2+1+1+X
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAs7h7d7c6c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAd6d7h7dXn6c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsXn7d7c6c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7d9h7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAhXn9h7c"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardFullHouseAndThreeKind(&alice) {
			//alice.Type = model.FOUROFAKIND
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardFlush(t *testing.T) {
	// 5, 4+X
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AdKd9d7h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAd6d7h7d3d5d"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsXn7d7c6c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7dXn7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAh7d9h7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdAsAhXn9h7c"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardFlush(&alice) {
			alice.Type = model.FlUSH
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardTwoPairs(t *testing.T) {
	//2+2+1
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AdKd9d7h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAd6d7h7d3d5d"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAdKsXn7dQc6c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcJd5sTh7dJc7c"})
	PokerList = append(PokerList, model.HandCards{Src: "Ac6d5sAh7d9h7c"})
	PokerList = append(PokerList, model.HandCards{Src: "AcAd8s8h9c9h7c"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardTwoPairs(&alice) {
			alice.Type = model.TWOPAIR
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardOnePairs(t *testing.T) {
	// 2+1+1+1, 1+1+1+1+X
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AdKd9d5h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "Ad4d9d5h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "AdTdJd8h7d7cKd"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAd6d7hTd3d5d"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAd6d6hTd3c3d"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardOnePairs(&alice) {
			alice.Type = model.ONEPAIR
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardNoPairs(t *testing.T) {
	// 1+1+1+1+1
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "AdKd9d5h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "Ad4d9d5h7d7c6d"})
	PokerList = append(PokerList, model.HandCards{Src: "AdTdJd8h7d7cKd"})
	PokerList = append(PokerList, model.HandCards{Src: "XnAd6d7hTd3d5d"})

	for k, alice := range PokerList {
		alice.IsGhost, _ = HasGhost(alice.Src)
		SortCard(&alice)
		if SevenCardNoPairs(&alice) {
			alice.Type = model.NOPAIR
		}
		fmt.Print(k+1, ". ", alice)
		fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	}
}

func TestSevenCardStraightFlushAndFlush(t *testing.T) {
	var PokerList []model.HandCards
	PokerList = append(PokerList, model.HandCards{Src: "6d7dXn3hTd5h4d"})
	//PokerList = append(PokerList, model.HandCards{Src:"Ad4d9d5h7d7c6d"})
	//PokerList = append(PokerList, model.HandCards{Src:"AdTdJd8h7d7cKd"})
	//PokerList = append(PokerList, model.HandCards{Src:"XnAd6d7hTd3d5d"})

	//for k, alice := range PokerList {
	//	alice.IsGhost, _ = HasGhost(alice.Src)
	//	SortCard(&alice)
	//	if SevenCardStraightFlushAndFlush(&alice, ) {
	//		alice.Type = model.NOPAIR
	//	}
	//	fmt.Print(k+1, ". ", alice)
	//	fmt.Println(" handcard type: ", model.HandCardType[alice.Type])
	//}
}

func TestParseTwoHandCard(t *testing.T) {
	alice := model.HandCards{Src: "QdQd8c4s8c"}
	bob := model.HandCards{Src: "Qc9sAcAsQh"}

	timer("", func() { ParseTwoHandCard(&alice, &bob) })
}

//=========================7张牌测试end=======================
