package service

/*
主要做分类功能,将所有卡牌组进行分类

*/

import (
	"strings"
	"texasPoker/model"
	//"regexp"
)

/*
*description: 判断是否为顺子
*param: card: 一手卡牌
*return: 判断结果,错误信息
 */
func CardIsStraight(cardFace string) bool {

	if _, ok := model.StraightList[cardFace]; ok {
		return true
	}

	return false
}

/*
*description: 判断花色是否都一样
*param: card: 一手卡牌的牌色
*return: 判断结果,错误信息
 */
func CardIsFlush(cardColor string) bool {
	if strings.Count(cardColor, cardColor[0:1]) == model.CardAmount {
		return true
	}
	return false
}

// 判断牌面是否为皇家同花顺, P.S.输入必须为排序后的同花牌面
func CardIsRoyalStraight(cardface string) bool {
	return model.STJQKA == cardface
}

// 同花的子分类
func FlushSubClassify(handcard *model.HandCards) {
	// 判断是否为同花
	if isStraight := CardIsStraight(handcard.SortColor); isStraight {
		//fmt.Printf("card %s 是同花顺, %s\n", handcard.Src, handcard.Sort)
		if IsRoyalStraight := CardIsRoyalStraight(handcard.SortFace); true == IsRoyalStraight {
			handcard.Type = model.ROYALFLUSH // 皇家同花顺
		} else {
			handcard.Type = model.STRAIGHTFLUSH // 同花顺
		}
	} else {
		handcard.Type = model.FlUSH // 同花
	}
	return
}

// 非同花牌面的子分类
func NotFlushSubClassify(handcard *model.HandCards) {
	// 判断是否为顺子
	if isStraight := CardIsStraight(handcard.SortColor); isStraight {
		handcard.Type = model.STRAIGHT // 顺子
	} else {
		OtherSubClassify(handcard)
	}
	return
}

//判断牌面是否为四条
func CardIsFourOfAKind(cardface string) bool {
	l := len(cardface)
	for i := 0; i < l; i++ {
		if 4 == strings.Count(cardface, cardface[i:i+1]) {
			return true
		}
		if i == 1 {
			return false
		}
	}
	return false
}

//是否包含Onepair
func CardContainOnePair(cardface string) bool {
	l := len(cardface)
	for i := 0; i < l; i++ {
		count := strings.Count(cardface, cardface[i:i+1])
		if 2 == count {
			return true
		} else if i == 3 {
			return false
		}
	}
	return false
}

//是否包含三条
func CardIsThreeOfAKind(cardface string) bool {
	l := len(cardface)
	for i := 0; i < l; i++ {
		face := cardface[i : i+1]
		count := strings.Count(cardface, face)
		if 3 == count {
			return true
		} else if i == 2 {
			return false
		}
	}
	return false
}

//是否是二对
func CardIsTwoPair(cardface string) bool {
	pairs := 0
	l := len(cardface)
	for i := 0; i < l-1; {
		count := strings.Count(cardface, cardface[i:i+1])
		if 2 == count {
			pairs++
			if 2 == pairs {
				return true
			}
		}
		i = i + count
	}

	return false
}

//是否包含一对
func CardIsOnePair(cardface string) bool {
	l := len(cardface)
	for i := 0; i < l-1; {
		count := strings.Count(cardface, cardface[i:i+1])
		if 2 == count {
			return true
		}
		i = i + count
	}
	return false
}

//其他牌类分类
func OtherSubClassify(handcard *model.HandCards) {
	//1. 四条
	if CardIsFourOfAKind(handcard.SortFace) {
		handcard.Type = model.FOUROFAKIND
		return
	}
	//2. 俘虏/三条
	if CardIsThreeOfAKind(handcard.SortFace) {
		if CardContainOnePair(handcard.SortFace) {
			handcard.Type = model.FULLHOUSE
		} else {
			handcard.Type = model.THREEOFAKIND
		}
		return
	}
	//3. 二对
	if CardIsTwoPair(handcard.SortFace) {
		handcard.Type = model.TWOPAIR
		return
	}
	//4. 一对
	if CardIsOnePair(handcard.SortFace) {
		handcard.Type = model.ONEPAIR
		return
	}
	//5. 单张
	handcard.Type = model.NOPAIR

	return
}

// 对输入的一手卡牌归类
func FiveCardParseV1(handcard *model.HandCards) {
	//判断是否为顺子
	if isStraight := CardIsStraight(handcard.SortFace); isStraight {
		FlushSubClassify(handcard)
	} else {
		NotFlushSubClassify(handcard)
	}

	return
}

func FiveCardParse(handcard *model.HandCards) {
	//判断是否为花
	if isFlush := CardIsFlush(handcard.SortFace); isFlush {
		FlushSubClassify(handcard)
	} else {
		NotFlushSubClassify(handcard)
	}

	return
}
