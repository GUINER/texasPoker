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



// 快速排列算法
func quickSort(CardList *[]model.CardFace, start int ,end int ) {

	var i,j = start,end

	//基值
	key := (*CardList)[start]

	for ; i < j;  {
		//不小于基值的
		for ; (*CardList)[j].Index >= key.Index && j > i; {
			j--
		}

		if j > i {
			//小于基准值的数往前扔
			(*CardList)[i] = (*CardList)[j]
			i++

			for ;(*CardList)[i].Index <= key.Index && i < j; {
				i++
			}
			if i < j {
				//大于基准值的往后扔
				(*CardList)[j] = (*CardList)[i]
				j--
			}
		}
	}

	(*CardList)[i] = key
	//fmt.Printf("排序中：%v,i=%d,j=%d\n", *CardList, i, j)
	if start < i {
		quickSort(CardList, start, i )
	}
	if j+1 < end {
		quickSort(CardList, j+1, end )
	}
}

/*
*description: 对一手卡牌进行排序,升序
*param: card: 一手卡牌
*return: 排序后的卡牌
*/
func SortCard(handcard *model.HandCards)  string {
	card := handcard.Src
	if "" == card || 0 ==len(card) {
		return ""
	}

	var CardList []model.CardFace


	// 解析牌面
	for i := 0; i < model.CardAmount; i++ {
		n := 2 * i
		m := n + 2

		letter := card[n : n+1]
		singleCard := model.CardFace{
			Face: letter,
			Color: card[n+1: m],
			Index: model.CardLetters[letter],
		}
		CardList = append(CardList, singleCard)
	}

	//fmt.Println("排序前：", CardList)
	// 对数组排序
	quickSort(&CardList,0, len(CardList) - 1)
	//fmt.Println("排序后：", CardList)

	//重新生成新的字符串
	newSort := ""
	for _,v := range CardList {
		newSort = newSort + v.Face + v.Color
		handcard.SortFace = handcard.SortFace + v.Face
		handcard.SortColor = handcard.SortColor + v.Color
	}
	handcard.Sort = newSort
	return newSort
}

/*
*description: 判断是否为顺子
*param: card: 一手卡牌
*return: 判断结果,错误信息
*/
func CardIsStraight(cardFace string) (bool, error) {
	if cardFace == "" || len(cardFace) != model.CardAmount {
		fmt.Errorf("the cardFace %s is invalod", cardFace)
		return false,fmt.Errorf("format is invalid")
	}

	//strs := make(map[int]string, model.CardAmount)
	//cardFace := ""
	//for i := 0; i < model.CardAmount ; i++ {
	//	n := 2 * i
	//	strs[i] = card[ n : n +1 ]
	//
	//	// 提取牌面信息
	//	cardFace = cardFace + strs[i]
	//}
	//fmt.Println(cardFace)

	// 牌面是否顺子
	if isExist := strings.Contains(model.CardStr, cardFace); isExist == false {
		return false, nil
	}

	return true,nil
}


/*
*description: 判断花色是否都一样
*param: card: 一手卡牌的牌色
*return: 判断结果,错误信息
*/
func CardIsFlush(cardColor string) (bool, error) {
	//cardColors := make([]string, model.CardAmount)
	//cardColorStr := ""
	//for i:= 0; i < model.CardAmount ; i++ {
	//	n := i*2 + 1
	//	//提取牌面的花色
	//	cardColors = append(cardColors, card[n:n+1])
	//	cardColorStr = cardColorStr + card[n:n+1]
	//}

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
		fmt.Errorf("CardIsRoyalStraight card[%s] is null ", cardface)
		return false
	}

	//reg, err := regexp.Compile("TJQKA")
	//if err != nil {
	//	fmt.Errorf("CardIsRoyalStraight card[%s] failed, error: %v", cardface, err)
	//	return false
	//}
	//
	//if false == reg.MatchString(cardface) {
	//	return false
	//}

	if "TJQKA" == cardface {
		//fmt.Errorf("CardIsRoyalStraight card[%s] failed", cardface)
		return true
	}
	return false
}

// 顺子子分类
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
			OtherSubClassify(handcard)
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
	for i := 0; i < len(cardface)-1 ;  {
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
		return nil
	}

	//2. 俘虏/三条
	if CardIsThreeOfAKind(handcard.SortFace) {
		if CardContainOnePair(handcard.SortFace) {
			handcard.Type = model.FULLHOUSE
		} else {
			handcard.Type = model.THREEOFAKIND
		}
		return nil
	}

	//3. 二对
	if CardIsTwoPair(handcard.SortFace) {
		handcard.Type = model.TWOPAIR
		return nil
	}

	//4. 一对
	if CardIsOnePair(handcard.SortFace) {
		handcard.Type = model.ONEPAIR
		return nil
	}
	//5. 单张
	handcard.Type = model.NOPAIR


	return fmt.Errorf("OtherSubClassify error: unknown card [%s]", handcard.Src)
}

// 对输入的一手卡牌归类
func ClassifyCard(handcard *model.HandCards) (err error) {

	//根据牌面排序
	if handcard.Sort = SortCard(handcard); "" == handcard.Sort {
		fmt.Errorf("SortCard %s failed", handcard.Src)
		return fmt.Errorf("SortCard failed")
	}

	//判断是否为顺子
	if IsStraight, err := CardIsStraight(handcard.SortFace); err != nil {
		fmt.Errorf("service.CardIsStraight error: %s", err.Error())
	} else {
		if IsStraight == true {	// 顺子类
			StraightSubClassify(handcard)
		} else {	//非顺子类
			//fmt.Printf("card %s 不是顺子, %s\n", handcard.Src, handcard.Sort)
			NotStraightSubClassify(handcard)
		}
	}

	return nil
}