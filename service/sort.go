package service

import (
	"texasPoker/model"
)

// 快速排列算法
func quickSort(CardList []*model.CardFace, start int, end int) {

	var i, j = start, end

	//基值
	key := CardList[start]

	for i < j {
		//不小于基值的
		for CardList[j].Index >= key.Index && j > i {
			j--
		}

		if j > i {
			//小于基准值的数往前扔
			CardList[i] = CardList[j]
			i++

			for CardList[i].Index <= key.Index && i < j {
				i++
			}
			if i < j {
				//大于基准值的往后扔
				CardList[j] = CardList[i]
				j--
			}
		}
	}

	CardList[i] = key
	//fmt.Printf("排序中：%v,i=%d,j=%d\n", *CardList, i, j)
	if start < i {
		quickSort(CardList, start, i)
	}
	if j+1 < end {
		quickSort(CardList, j+1, end)
	}
}

// 第三种写法
func quick3Sort(CardList []*model.CardFace, left int, right int) {

	if left >= right {
		return
	}

	explodeIndex := left

	for i := left + 1; i <= right; i++ {

		if CardList[left].Index >= CardList[i].Index {

			//分割位定位++
			explodeIndex++
			CardList[i], CardList[explodeIndex] = CardList[explodeIndex], CardList[i]

		}

	}

	//起始位和分割位
	CardList[left], CardList[explodeIndex] = CardList[explodeIndex], CardList[left]

	quick3Sort(CardList, left, explodeIndex-1)
	quick3Sort(CardList, explodeIndex+1, right)
}

/*
*description: 对一手卡牌进行排序,升序
*param: card: 一手卡牌
*return: 排序后的卡牌
 */
func SortCard(handcard *model.HandCards) {
	var CardList []*model.CardFace
	card := handcard.Src

	// 解析牌面
	length := len(card)
	for i := 0; i < length; i += 2 {
		m := i + 1
		letter := card[i:m]
		CardList = append(CardList, &model.CardFace{
			Face:  letter,
			Color: card[m : m+1],
			Index: model.CardLetters[letter],
		})
	}

	//fmt.Printf("排序前：%v\n", CardList)
	quickSort(CardList, 0, len(CardList)-1)
	//fmt.Printf("排序后：%v\n", CardList)

	for _, v := range CardList {
		handcard.Sort = handcard.Sort + v.Face + v.Color
		handcard.SortFace = handcard.SortFace + v.Face
		handcard.SortColor = handcard.SortColor + v.Color
	}
}

func SortCardV1(p *model.HandCards) {
	var pokers []*model.Poker
	card := p.Src

	length := len(card)
	for i := 0; i < length; i += 2 {

		pokers = append(pokers, &model.Poker{
			Letter: model.CardLetters[card[i:i]],
			Color:  card[i+1 : i+2],
		})
	}
}
