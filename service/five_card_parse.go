package service

/*
该业务成主要做分类功能,将所有卡牌组进行分类

*/

import (
	"fmt"
	"myproject/texasPoker/model"
	"strings"
	//"regexp"
)


/*
*description: 判断是否为顺子
*param: card: 一手卡牌
*return: 判断结果,错误信息
*/
func CardIsStraight(cardFace string) (bool, error) {

	if cardFace == ""  {
		return false, fmt.Errorf("the cardFace %s is invalid", cardFace)
	}
	lenght := len(cardFace)
	if model.FiveCard == lenght {
		idx,ok := model.StraightList[cardFace]
		if !ok && idx < 1 {
			return false, nil
		}
	}

	return true,nil
}


/*
*description: 判断花色是否都一样
*param: card: 一手卡牌的牌色
*return: 判断结果,错误信息
*/
func CardIsFlush(cardColor string) (bool, error) {

	for _,v := range model.CardColor {
		if strings.Count(cardColor, v) >= model.CardAmount {
			return true, nil
		}
	}

	return false, nil
}

// 判断牌面是否为皇家同花顺, P.S.输入必须为排序后的同花牌面
func CardIsRoyalStraight(cardface string) bool {

	if "" == cardface {
		fmt.Printf("CardIsRoyalStraight card[%s] is null ", cardface)
		return false
	}

	if model.STJQKA == cardface {
		return true
	}
	return false
}

// 顺子的子分类
func StraightSubClassify(handcard *model.HandCards) (err error) {
	// 判断是否为同花
	if IsFlush,err := CardIsFlush(handcard.SortColor); err != nil {
		return fmt.Errorf("service.CardIsFlush error: %s", err.Error())
	} else {
		if true == IsFlush {
			//fmt.Printf("card %s 是同花顺, %s\n", handcard.Src, handcard.Sort)
			if IsRoyalStraight := CardIsRoyalStraight(handcard.SortFace); true == IsRoyalStraight {
				handcard.Type = model.ROYALFLUSH	// 皇家同花顺
			} else {
				handcard.Type = model.STRAIGHTFLUSH	// 同花顺
			}
		} else {
			handcard.Type = model.STRAIGHT	//顺子
		}
	}

	return nil
}

// 非顺子牌面的子分类
func NotStraightSubClassify(handcard *model.HandCards) (err error) {
	// 判断是否为同花
	if IsFlush,err := CardIsFlush(handcard.SortColor); err != nil {
		return fmt.Errorf("service.CardIsFlush error: %s", err.Error())
	} else {
		if true == IsFlush {
			handcard.Type = model.FlUSH	// 同花
		} else {
			err = OtherSubClassify(handcard)
		}
	}
	return nil
}


//判断牌面是否为四条
func CardIsFourOfAKind(cardface string) (bool) {
	for _,v := range model.FaceList {
		if 4 == strings.Count(cardface, v) {
			return true
		}
	}
	return false
}

//是否包含Onepair
func CardContainOnePair(cardface string) (bool) {
	for i := 0 ; i < len(cardface) ; i++ {
		face := cardface[i:i+1]
		count := strings.Count(cardface, face)
		if 2 == count {
			return true
		}
	}
	return false
}

//是否包含三条
func CardIsThreeOfAKind(cardface string) (bool) {
	for i := 0 ; i < len(cardface) ; i++ {
		face := cardface[i:i+1]
		count := strings.Count(cardface, face)
		if 3 == count {
			return true
		}
	}
	return false
}

//是否是二对
func CardIsTwoPair(cardface string) bool {
	pairs := 0
	for i := 0; i < len(cardface)-1 ;  {
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
	for i := 0; i < len(cardface)-1 ; {
		count := strings.Count(cardface, cardface[i:i+1])
		if 2 == count {
			return true
		}
		i = i + count
	}
	return false
}

//其他牌类分类
func OtherSubClassify(handcard *model.HandCards) (err error) {
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
func FiveCardParse(handcard *model.HandCards) (err error) {
	//判断是否为顺子
	if IsStraight, err := CardIsStraight(handcard.SortFace); err != nil {
		return fmt.Errorf("CardIsStraight error: %s", err.Error())
	} else {
		if IsStraight == true {	// 顺子类
			err = StraightSubClassify(handcard)
		} else {	//非顺子类
			err = NotStraightSubClassify(handcard)
		}
	}

	return nil
}