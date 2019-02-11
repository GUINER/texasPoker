package service

import (
	"fmt"
	"myproject/texasPoker/model"
	"strings"
)

func SevenCardHasStraight(cardFace string) (bool) {
	for i := 0; i < len(cardFace); i++ {
		letter := cardFace[i:i+1]
		count := strings.Count(cardFace, letter)
		if count > 1 {
			cardFace = strings.Replace(cardFace, letter,"", count-1)
		}
	}
	fmt.Println(cardFace)
	if len(cardFace) < model.CardAmount {
		return false
	}
	idx,ok := model.StraightList[cardFace]
	if !ok && idx < 1 {
		return false
	}
	return true
}

func SevenCardParse(alice,bob *model.HandCards) {

}

func ParseTwoHandCard(alice,bob *model.HandCards) {

	//1. 5张
	if model.FiveCard == len(alice.SortFace) {
		//不需要任何处理
		return
	}
	//2. 5张带赖子

	//3. 7张
	if model.SevenCard == len(alice.SortFace) {
		//
		SevenCardParse(alice,bob)
		return
	}
	//4. 7张带赖子

}