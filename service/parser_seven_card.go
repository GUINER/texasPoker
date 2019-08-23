package service

import (
	"fmt"
	"strings"
	"texasPoker/model"
)

// 7张牌中是否有顺子
func SevenCardHasStraight(cardFace string) bool {
	for i := 0; i < len(cardFace); i++ {
		letter := cardFace[i : i+1]
		count := strings.Count(cardFace, letter)
		if count > 1 {
			cardFace = strings.Replace(cardFace, letter, "", count-1)
		}
	}
	fmt.Println(cardFace)
	if len(cardFace) < model.CardAmount {
		return false
	}
	// 顺子
	_, ok := model.StraightList[cardFace]
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

// 是否包含顺子
func SevenCardHasStraight2(any *model.HandCards) (bool, string) {
	countLimit := 5

	if any.IsGhost {
		countLimit = 4
	}

	for j := len(model.StraightRankList) - 1; j >= 0; j-- {
		count := 0
		straight := model.StraightRankList[j]
		tmpstr := any.SortFace
		for n := len(straight) - 1; n >= 0; n-- {
			if strings.Contains(tmpstr, straight[n:n+1]) {
				tmpstr = strings.Replace(tmpstr, straight[n:n+1], "", -1)
				count++
			}
		}
		if count >= countLimit {
			any.Type = model.STRAIGHT
			return true, straight
		}
	}

	return false, ""
}

// 区分同花顺和同花
func SevenCardStraightFlushAndFlush(any *model.HandCards, s string) {
	countLimit := 5
	if any.IsGhost {
		countLimit = 4
	}

	for _, color := range model.CardColor {
		if count := strings.Count(any.Sort, color); count >= countLimit { //花色
			tmpface := ""
			matchList := []string{}
			handcard := any.Sort
			for i := 0; i < count; i++ {
				index := strings.IndexAny(handcard, color)
				face := handcard[index-1 : index]
				matchList = append(matchList, face)
				handcard = strings.Replace(handcard, face+color, "", 1)
				tmpface = tmpface + face
				//fmt.Println(tmpface)
			}
			if any.IsGhost {
				tmpface = tmpface + model.Ghost
			}

			count = 0
			for _, v := range matchList {
				if strings.Contains(s, v) {
					count++
				}
			}

			if count >= countLimit && model.STRAIGHT == any.Type {
				any.SortFace = s
				any.Type = model.STRAIGHTFLUSH
			} else {
				any.SortFace = tmpface
				any.Type = model.FlUSH
			}
			return
		}
	}
	if model.STRAIGHT == any.Type {
		any.SortFace = s
		any.Type = model.STRAIGHT
	}
	return
}

// 四条
func SevenCardFourOfAKind(any *model.HandCards) bool {
	//4+X, 4+1, 3+1+X
	cardface := any.SortFace
	lenght := len(cardface)

	numLimit := 4
	if any.IsGhost {
		numLimit = 3
		lenght--
	}

	for i := lenght - 1; i >= 0; i-- {
		face := cardface[i : i+1]
		if count := strings.Count(cardface, face); count >= numLimit {
			// 属于4条
			any.SortFace = face + face + face + face
			if 4 == count {
				if any.IsGhost { // 4+X
					any.SortFace = any.SortFace + model.Ghost
				} else { // 4+1
					if i == lenght-1 {
						any.SortFace = cardface[i-4:i-3] + any.SortFace
					} else if i < lenght-1 {
						any.SortFace = any.SortFace + cardface[lenght-1:lenght]
					}
				}
			} else {
				// 3+1+X
				cardface = strings.Replace(cardface, face, "", count)
				cardface = strings.Replace(cardface, model.Ghost, "", 1)
				lenght = len(cardface)
				b := cardface[lenght-1 : lenght]
				result := compareLetter(face, b)
				if model.GREAT == result {
					any.SortFace = b + any.SortFace
				} else if model.LESS == result {
					any.SortFace = any.SortFace + b
				}
			}

			return true
		}
	}

	return false
}

func genStrbyResult(a, b, c string, result int) string {
	if model.GREAT == result {
		return b + c + a
	} else {
		return a + b + c
	}
}

// 俘虏/3条
func SevenCardFullHouseAndThreeKind(any *model.HandCards) bool {
	// 2+2+X, 3+2 or 3+1+1, 2+1+1+X
	cardface := any.SortFace
	lenght := len(cardface)
	numLimit := 3

	if any.IsGhost {
		numLimit = 2
		lenght--
	}

	for i := lenght - 1; i >= 0; i-- {
		face := cardface[i : i+1]
		if count := strings.Count(cardface, face); count >= numLimit {
			any.SortFace = face + face + face
			cardface = strings.Replace(cardface, face, "", count)
			lenght = len(cardface)
			letter := cardface[lenght-1 : lenght]

			result := compareLetter(face, letter)
			if 3 == count {
				if count := strings.Count(cardface, letter); 2 == count { // 3+2
					any.SortFace = genStrbyResult(any.SortFace, letter, letter, result)
					any.Type = model.FULLHOUSE
				} else { // 3+1+1
					secondLetter := cardface[lenght-2 : lenght-1]
					if model.GREAT == result {
						any.SortFace = secondLetter + letter + any.SortFace
					} else {
						if model.GREAT == compareLetter(face, secondLetter) {
							any.SortFace = secondLetter + any.SortFace + letter
						} else {
							any.SortFace = any.SortFace + secondLetter + letter
						}
					}
					any.Type = model.THREEOFAKIND
				}
			} else { // 2+2+X
				if count := strings.Count(cardface, letter); 2 == count {
					any.SortFace = genStrbyResult(any.SortFace, letter, letter, result)
					any.Type = model.FULLHOUSE
				} else { // 2+1+1+X
					secondLetter := cardface[lenght-2 : lenght-1]
					if model.GREAT == result {
						any.SortFace = secondLetter + letter + any.SortFace
					} else {
						if model.GREAT == compareLetter(face, secondLetter) {
							any.SortFace = secondLetter + any.SortFace + letter
						} else {
							any.SortFace = any.SortFace + secondLetter + letter
						}
					}
					any.Type = model.THREEOFAKIND
				}
			}
			return true
		}
	}
	return false
}

// 同花
func SevenCardFlush(any *model.HandCards) bool {
	// 5, 4+X
	countLimit := 5

	if any.IsGhost {
		countLimit = 4
	}
	for _, color := range model.CardColor {
		if count := strings.Count(any.SortColor, color); count >= countLimit {
			cardface := any.SortFace
			lenght := len(cardface)
			tmpface := ""
			num := 0

			for i := lenght - 1; i >= 0; i-- {
				tmpcolor := any.SortColor[i : i+1]
				if color == tmpcolor || "n" == tmpcolor {
					tmpface = any.SortFace[i:i+1] + tmpface
					num++
					if 5 == num {
						any.SortFace = tmpface
						return true
					}
				}
			}

			return false
		}
	}
	return false
}

func SevenCardThreeOfAkind(any *model.HandCards) bool {
	// 3+1+1, 2+1+1+X
	return false
}

func SevenCardTwoPairs(any *model.HandCards) bool {
	//2+2+1
	if any.IsGhost {
		return false
	}
	countLimit := 2
	pairs := 0
	single := 0
	cardface := any.SortFace
	lenght := len(cardface)
	tmpface := ""

	for i := lenght - 1; i >= 0; i-- {
		face := cardface[i : i+1]

		if count := strings.Count(cardface, face); count == countLimit && pairs != countLimit {
			if pairs == countLimit {
				continue
			}
			i--
			tmpface = face + face + tmpface
			pairs++
		} else if 0 == single {
			tmpface = face + tmpface
			single++
		}
	}
	if 1 == single && countLimit == pairs {
		any.SortFace = tmpface
		return true
	}

	return false
}

// 一对
func SevenCardOnePairs(any *model.HandCards) bool {
	//2+1+1+1, 1+1+1+1+X
	cardface := any.SortFace
	lenght := len(cardface)

	if any.IsGhost {
		//fmt.Println(cardface, any)
		cardface = strings.Replace(cardface, model.Ghost, cardface[lenght-2:lenght-1], 1)
		any.SortFace = cardface[lenght-5 : lenght]
		//fmt.Println(any.SortFace)
		return true
	}
	countLimit := 2
	single, pairs := 0, 0
	tmpface := ""

	for i := lenght - 1; i >= 0; i-- {
		face := cardface[i : i+1]
		if count := strings.Count(cardface, face); count == countLimit {
			if 1 == pairs {
				continue
			}
			tmpface = face + face + tmpface
			pairs++
		} else if single < 3 && 1 == count {
			tmpface = face + tmpface
			single++
		}
		if 3 == single && 1 == pairs { //2+1+1+1
			any.SortFace = tmpface
			return true
		}
	}
	return false
}

// 单张
func SevenCardNoPairs(any *model.HandCards) bool {
	//1+1+1+1+1
	if any.IsGhost {
		return false
	}
	cardface := any.SortFace
	lenght := len(cardface)
	any.SortFace = cardface[lenght-5 : lenght]

	return true
}

// 其他牌型解析
func SevenCardOtherClassParse(any *model.HandCards) {
	// 四条、俘虏、同花、三条、二对、一对、单张
	if SevenCardFourOfAKind(any) {
		any.Type = model.FOUROFAKIND
		return
	}
	if SevenCardFullHouseAndThreeKind(any) {
		return
	}
	//if SevenCardFlush(any) {
	//	any.Type = model.FlUSH
	//	return
	//}
	//if SevenCardThreeOfAkind(any) {
	//
	//}
	if SevenCardTwoPairs(any) {
		any.Type = model.TWOPAIR
		return
	}
	if SevenCardOnePairs(any) {
		any.Type = model.ONEPAIR
		return
	}
	if SevenCardNoPairs(any) {
		any.Type = model.NOPAIR
		return
	}
	any.Type = model.UNKNOWN
	return
}

//	解析7张牌
func SevenCardParse(any *model.HandCards) {
	// 是否包含顺子
	if isStraight, mStraight := SevenCardHasStraight2(any); isStraight {
		// 是否包含同花
		if SevenCardHasFlush(any) {
			// 分出同花顺、同花
			//fmt.Println(any.Src)
			SevenCardStraightFlushAndFlush(any, mStraight)
		} else {
			any.SortFace = mStraight
		}
	} else { // 没有顺子
		SevenCardOtherClassParse(any)
	}
}
