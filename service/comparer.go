package service

import (
	"fmt"
	"texasPoker/model"
)

type PokerComparer struct {
	Alice  *model.HandCards
	Bob    *model.HandCards
	Result int
}

func NewPokerComparer(alice, bob string) *PokerComparer {
	return &PokerComparer{
		Alice: &model.HandCards{Src: alice},
		Bob:   &model.HandCards{Src: bob},
	}
}

func (pc *PokerComparer) sortCard() {
	SortCard(pc.Alice)
	SortCard(pc.Bob)
}

func (pc *PokerComparer) parseCard() {
	ParseTwoHandCard(pc.Alice, pc.Bob)
}

func (pc *PokerComparer) compare() {
	pc.Result = compareByType(pc.Alice.Type, pc.Bob.Type)

	if model.EQUAL == pc.Result {
		pc.Result = compareByFace(pc.Alice.SortFace, pc.Bob.SortFace, pc.Alice.Type)
	}
}

func (pc *PokerComparer) Compare() {
	pc.sortCard()

	pc.parseCard()

	pc.compare()
}

func (pc *PokerComparer) PrintResult() {
	switch {
	case pc.Result == model.EQUAL:
		fmt.Printf("result: alice[%s][%s] = bob[%s][%s]\n", pc.Alice.Src, model.HandCardType[pc.Alice.Type], pc.Bob.Src, model.HandCardType[pc.Bob.Type])
	case pc.Result == model.GREAT:
		fmt.Printf("result: alice[%s][%s] > bob[%s][%s]\n", pc.Alice.Src, model.HandCardType[pc.Alice.Type], pc.Bob.Src, model.HandCardType[pc.Bob.Type])
	case pc.Result == model.LESS:
		fmt.Printf("result: alice[%s][%s] < bob[%s][%s]\n", pc.Alice.Src, model.HandCardType[pc.Alice.Type], pc.Bob.Src, model.HandCardType[pc.Bob.Type])
	default:
		fmt.Printf("unknown result[%d]: alice[%s][%s] - bob[%s][%s]\n", pc.Result, pc.Alice.Src, model.HandCardType[pc.Alice.Type], pc.Bob.Src, model.HandCardType[pc.Bob.Type])
	}
}
