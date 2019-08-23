package service

import (
	"texasPoker/model"
)

// 快速排列算法
func quickSort(CardList *[]model.CardFace, start int, end int) {

	var i, j = start, end

	//基值
	key := (*CardList)[start]

	for i < j {
		//不小于基值的
		for (*CardList)[j].Index >= key.Index && j > i {
			j--
		}

		if j > i {
			//小于基准值的数往前扔
			(*CardList)[i] = (*CardList)[j]
			i++

			for (*CardList)[i].Index <= key.Index && i < j {
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
		quickSort(CardList, start, i)
	}
	if j+1 < end {
		quickSort(CardList, j+1, end)
	}
}

/*
*description: 对一手卡牌进行排序,升序
*param: card: 一手卡牌
*return: 排序后的卡牌
 */
func SortCard(handcard *model.HandCards) {
	var CardList []model.CardFace
	card := handcard.Src

	// 解析牌面
	for i := 0; i < len(card)/2; i++ {
		n := 2 * i
		m := n + 2

		letter := card[n : n+1]
		CardList = append(CardList, model.CardFace{
			Face:  letter,
			Color: card[n+1 : m],
			Index: model.CardLetters[letter],
		})
	}

	//fmt.Println("排序前：", CardList)
	quickSort(&CardList, 0, len(CardList)-1)
	//fmt.Println("排序后：", CardList)

	handcard.Sort = ""
	for _, v := range CardList {
		handcard.Sort = handcard.Sort + v.Face + v.Color
		handcard.SortFace = handcard.SortFace + v.Face
		handcard.SortColor = handcard.SortColor + v.Color
	}
}
