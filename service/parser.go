package service

import (
	"fmt"
	"myproject/texasPoker/model"
	"strings"
)

// 7张牌中是否有顺子
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
	// 顺子
	_,ok := model.StraightList[cardFace]
	if !ok {
		return false
	}
	return true
}

// 是否有赖子
func HasGhost(cardface string) (bool, int) {
	if count := strings.Count(cardface, model.Ghost); count > 0 {
		return true, count
	}
	return false, 0
}

// 判断同花
func SevenCardHasFlush(any *model.HandCards) bool {
	countLimit := 5

	if any.IsGhost {
		countLimit = 4
	}
	for _, v := range model.CardColor {
		if count := strings.Count(any.SortColor, v); count >= countLimit {
			return true
		}
	}
	return false
}

// 是否包含赖子
func SevenCardHasStraight2(any *model.HandCards) (bool, string) {
	countLimit := 5

	if any.IsGhost {
		countLimit = 4
	}
	//var matchList []string
	for j := len(model.StraightRankList) - 1; j >= 0; j-- {
		count := 0
		straight := model.StraightRankList[j]
		tmpstr := any.SortFace
		for n := len(straight) - 1; n >= 0; n-- {
			if strings.Contains(tmpstr, straight[n:n+1]) {
				//fmt.Println("match string: ", tmpstr)
				tmpstr = strings.Replace(tmpstr, straight[n:n+1], "", -1)
				count++
			}
		}
		if count >= countLimit {
			any.Type = model.STRAIGHT
			fmt.Println("match straight: ", straight)
			return true, straight
		}
	}

	return false, ""
}

//	解析7张牌
func SevenCardParse(any *model.HandCards)  {

	 if isStraight, mStraight := SevenCardHasStraight2(any); isStraight {
		fmt.Println(mStraight)
	 }
}

func ParseTwoHandCard(alice, bob *model.HandCards) {
	// 是否带有赖子
	alice.IsGhost, _ = HasGhost(alice.SortFace)
	bob.IsGhost, _ 	 = HasGhost(bob.SortFace)

	//1. 5张
	if model.FiveCard == len(alice.SortFace) {
		return
	}
	//2. 7张
	if model.SevenCard == len(alice.SortFace) {
		SevenCardParse(alice)
		SevenCardParse(bob)
		return
	}

}