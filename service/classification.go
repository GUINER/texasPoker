package service

/*
该业务成主要做分类功能,将所有卡牌组进行分类

*/

import (
	"fmt"
	"myproject/texasPoker/model"
	"strings"
)



// 快速排列算法
func quickSort(CardList *[]model.CardFace, start int ,end int ) {

	//基值
	key_i := start
	key := (*CardList)[key_i]
	var i,j = start,end

	for ; i < j;  {
		//不小于基值的
		for ; (*CardList)[j].Index >= key.Index && j > i; {
			j--
		}

		if j > i {
			//大于基准值的往后扔
			(*CardList)[i] = (*CardList)[j]
			i++

			for ;(*CardList)[i].Index <= key.Index && i < j; {
				i++
			}
			if i < j {
				//小于基准值的数往前扔
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
func SortCard(card string)  string {

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
			//Number: i,
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
	}

	return newSort
}

/*
*description: 判断是否为顺子
*param: card: 一手卡牌
*return: 判断结果,错误信息
*/
func CardIsStraight(card string) (bool, error) {
	if card == "" || len(card) != model.AmountOfCardLetter {
		fmt.Errorf("the card %s is invalod", card)
		return false,fmt.Errorf("format is invalid")
	}

	strs := make(map[int]string, model.CardAmount)
	cardFace := ""
	for i := 0; i < model.CardAmount ; i++ {
		n := 2 * i
		strs[i] = card[ n : n +1 ]

		// 提取牌面信息
		cardFace = cardFace + strs[i]
	}
	//fmt.Println(cardFace)

	// 牌面是否顺子
	if isExist := strings.Contains(model.CardStr, cardFace); isExist == false {
		return false, nil
	}

	return true,nil
}


/*
*description: 判断花色是否都一样
*param: card: 一手卡牌
*return: 判断结果,错误信息
*/
func CardIsFlush(card string) (bool, error) {
	cardColors := make([]string, model.CardAmount)
	cardColorStr := ""
	for i:= 0; i < model.CardAmount ; i++ {
		n := i*2 + 1
		//提取牌面的花色
		cardColors = append(cardColors, card[n:n+1])
		cardColorStr = cardColorStr + card[n:n+1]
	}

	for _,v := range model.CardColor {
		if strings.Count(cardColorStr, v) >= model.CardAmount {
			return true, nil
		}
	}

	return false, nil
}




// 对输入的一手卡牌归类
func ClassifyCard(handcard *model.HandCards) (err error) {

	//handcard.Src = card1

	//根据牌面排序
	handcard.Sort = SortCard(handcard.Src)

	//判断是否为顺子
	if IsStraight, err := CardIsStraight(handcard.Sort); err != nil {
		fmt.Errorf("service.CardIsStraight error: %s", err.Error())
	} else {
		if IsStraight == true {
			fmt.Printf("card %s 是顺子, %s\n", handcard.Src, handcard.Sort)

			// 判断是否为同花
			if IsFlush,err := CardIsFlush(handcard.Sort); err != nil {
				fmt.Errorf("service.CardIsFlush error: %s", err.Error())
			} else {
				if true == IsFlush {
					fmt.Printf("card %s 是同花, %s\n", handcard.Src, handcard.Sort)
				}
			}
		} else {
			fmt.Printf("card %s 不是顺子, %s\n", handcard.Src, handcard.Sort)
		}
	}

	return nil
}