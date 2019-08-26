package service

import (
	"fmt"
	"texasPoker/model"
)

type PokerDealer struct {
	Alice  *model.HandCards
	Bob    *model.HandCards
	Result int
}

func NewPokerDealer(alice, bob string) *PokerDealer {
	return &PokerDealer{
		Alice: &model.HandCards{Src: alice},
		Bob:   &model.HandCards{Src: bob},
	}
}

func (pd *PokerDealer) Compare() {
	debugModel := true
	if debugModel {
		timer("sort card", func() { pd.sortCard() })
		timer("parse card", func() { pd.parseCard() })
		timer("compare card", func() { pd.compare() })
	} else {
		pd.sortCard()
		pd.parseCard()
		pd.compare()
	}
}

func (pd *PokerDealer) sortCard() {
	SortCard(pd.Alice)
	SortCard(pd.Bob)
}

func (pd *PokerDealer) parseCard() {
	ParseTwoHandCard(pd.Alice, pd.Bob)
}

func (pd *PokerDealer) compare() {
	if pd.Result = compareByType(pd.Alice.Type, pd.Bob.Type); model.EQUAL == pd.Result {
		pd.Result = compareByFace(pd.Alice.SortFace, pd.Bob.SortFace, pd.Alice.Type)
	}
}

func (pd *PokerDealer) PrintResult() {
	symbol := ""
	switch {
	case pd.Result == model.EQUAL:
		symbol = "="
	case pd.Result == model.GREAT:
		symbol = ">"
	case pd.Result == model.LESS:
		symbol = ">"
	default:
		symbol = "xxxxx" // unknown
	}
	fmt.Printf("result: alice[%s][%s] %s bob[%s][%s]\n", pd.Alice.Sort, model.HandCardType[pd.Alice.Type], symbol, pd.Bob.Sort, model.HandCardType[pd.Bob.Type])
}
