package service

import (
	"myproject/texasPoker/model"
	"strings"
)

func compareLetter(astr,bstr string) (result int) {
	a := model.CardLetters[astr]
	b := model.CardLetters[bstr]
	if a > b {
		return model.GREAT
	}
	if a < b {
		return model.LESS
	}
	return model.EQUAL
}


//单张
func CompareNoPair(aFace,bFace string) (result int) {
	len := len(aFace)

	for i := len; i > 0 ; i--{
		a := aFace[i - 1:i]
		b := bFace[i - 1:i]

		result = compareLetter(a,b)
		if model.EQUAL != result {
			return result
		}
	}


	return model.EQUAL
}

//一对
func CompareOnePair(aFace,bFace string) (result int) {
	len := len(aFace)
	for i := len; i > 0; i-- {
		a := aFace[i - 1:i]
		b := bFace[i - 1:i]

		result = compareLetter(a, b)
		if model.EQUAL != result {
			return result
		}
		if i == len {
			i--
		}
	}
	return model.EQUAL
}

//二对
func CompareTwoPair(aFace,bFace string) (result int) {
	len := len(aFace)
	for i := len; i > 0;  {
		a := aFace[i - 1:i]
		b := bFace[i - 1:i]

		result = compareLetter(a, b)
		if model.EQUAL != result {
			return result
		}
		i = i - 2
	}
	return model.EQUAL
}


//三条
func CompareThreeOfAKind(aFace,bFace string) (result int) {
	var a,b string
	lenght := len(aFace)
	for i := 0; i < lenght - 3; i++ {
		letter := aFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < lenght - 3 ; i++ {
		letter := bFace[i:i+1]
		if 3 == strings.Count(bFace, letter) {
			b = letter
			break
		}
	}
	result = compareLetter(a,b)
	if model.EQUAL == result {
		strings.Replace(aFace, a,"",3)
		strings.Replace(bFace, b,"",3)
		lenght = len(aFace)
		for i := lenght; i > 0; i-- {
			result = compareLetter(aFace[i-1:i], bFace[i-1:i])
			if model.EQUAL != result {
				return result
			}
		}
		return model.EQUAL
	}
	return compareLetter(a,b)
}


//顺子
func CompareStraight(aFace,bFace string) (result int) {
	Aidx,ok := model.StraightList[aFace]
	if !ok {
		return -1
	}
	Bidx,ok := model.StraightList[bFace]
	if !ok {
		return -1
	}
	if Aidx > Bidx {
		return model.GREAT
	}
	if Aidx < Bidx {
		return model.LESS
	}
	return model.EQUAL
}

//同花
func CompareFlush(aFace,bFace string) (result int) {
	len := len(aFace)
	a := aFace[len - 1:len]
	b := bFace[len - 1:len]

	return compareLetter(a,b)
}

//俘虏
func CompareFullHouse(aFace,bFace string) (result int) {
	var a,b string
	lenght := len(aFace)
	for i := 0; i < lenght - 3; i++ {
		letter := aFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < lenght - 3 ; i++ {
		letter := bFace[i:i+1]
		if 3 == strings.Count(bFace, letter) {
			b = letter
			break
		}
	}

	result = compareLetter(a,b)
	if model.EQUAL == result {
		strings.Replace(aFace, a,"",3)
		strings.Replace(bFace, b,"",3)
		lenght = len(aFace)
		//for i := lenght; i > 0; i-- {
			result = compareLetter(aFace[0:1], bFace[0:1])
			if model.EQUAL != result {
				return result
			}
		//}
		return model.EQUAL
	}

	return compareLetter(a,b)
}

//四条
func CompareFourOfAKind(aFace,bFace string) (result int) {
	var a,b string
	lenght := len(aFace)
	for i := 0; i < lenght - 3; i++ {
		letter := aFace[i:i+1]
		if 4 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < lenght - 3 ; i++ {
		letter := bFace[i:i+1]
		if 4 == strings.Count(bFace, letter) {
			b = letter
			break
		}
	}

	result = compareLetter(a,b)
	if model.EQUAL == result {
		strings.Replace(aFace,a,"",4)
		strings.Replace(bFace,b,"",4)
		lenght = len(aFace)
		for i := lenght; i > 0; i-- {
			result = compareLetter(aFace[i-1:i], bFace[i-1:i])
			if model.EQUAL != result {
				return result
			}
		}
		return model.EQUAL
	}

	return result
}

//同花顺
func CompareStraightFlush(aFace,bFace string) (result int) {
	Aidx,ok := model.StraightList[aFace]
	if !ok {
		return -1
	}
	Bidx,ok := model.StraightList[bFace]
	if !ok {
		return -1
	}
	if Aidx > Bidx {
		return model.GREAT
	}
	if Aidx < Bidx {
		return model.LESS
	}
	return model.EQUAL
}

//通过牌面face比较大小
func compareByFace(aliceFace, bobFace string, pokerType int) (result int) {

	switch pokerType {
	case model.ROYALFLUSH:
		return model.EQUAL
	case model.STRAIGHTFLUSH:
		return CompareStraightFlush(aliceFace, bobFace)
	case model.FOUROFAKIND:
		return CompareFourOfAKind(aliceFace, bobFace)
	case model.FULLHOUSE:
		return CompareFullHouse(aliceFace, bobFace)
	case model.FlUSH:
		return CompareFlush(aliceFace, bobFace)
	case model.STRAIGHT:
		return CompareStraight(aliceFace, bobFace)
	case model.THREEOFAKIND:
		return CompareThreeOfAKind(aliceFace, bobFace)
	case model.TWOPAIR:
		return CompareTwoPair(aliceFace, bobFace)
	case model.ONEPAIR:
		return CompareOnePair(aliceFace, bobFace)
	case model.NOPAIR:
		return CompareNoPair(aliceFace, bobFace)
	}
	return model.EQUAL
}

// 通过牌面的类型比较大小
func compareByType(alice, bob int) (result int) {
	if alice > bob {
		return model.GREAT
	}
	if alice < bob {
		return model.LESS
	}
	return model.EQUAL
}

func CompareTwoHandCard(Alice, Bob *model.HandCards) (result int) {

	//1. 根据牌面排序
	SortTwoHandCard(Alice, Bob)

	//2. 解析
	ParseTwoHandCard(Alice, Bob)

	//3. 分类
	ClassifyCard(Alice)
	ClassifyCard(Bob)

	//4. 比较大小
	result = compareByType(Alice.Type, Bob.Type)
	if model.EQUAL == result {
		//同一类型牌面比较
		return compareByFace(Alice.SortFace, Bob.SortFace, Alice.Type)
	}

	return result
}
