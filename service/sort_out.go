package service

import (
	"fmt"
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
	a := aFace[len - 1:len]
	b := bFace[len - 1:len]

	return compareLetter(a,b)
}

//三条
func CompareThreeOfAKind(aFace,bFace string) (result int) {
	var a,b string
	len := len(aFace)
	for i := 0; i < len - 3; i++ {
		letter := aFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < len - 3 ; i++ {
		letter := bFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			b = letter
			break
		}
	}
	return compareLetter(a,b)
}


//同花
func CompareStraight(aFace,bFace string) (result int) {
	len := len(aFace)
	a := aFace[len - 1:len]
	b := bFace[len - 1:len]

	return compareLetter(a,b)
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
	len := len(aFace)
	for i := 0; i < len - 3; i++ {
		letter := aFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < len - 3 ; i++ {
		letter := bFace[i:i+1]
		if 3 == strings.Count(aFace, letter) {
			b = letter
			break
		}
	}

	return compareLetter(a,b)
}

//四条
func CompareFourOfAKind(aFace,bFace string) (result int) {
	var a,b string
	len := len(aFace)
	for i := 0; i < len - 3; i++ {
		letter := aFace[i:i+1]
		if 4 == strings.Count(aFace, letter) {
			a = letter
			break
		}
	}
	for i := 0; i < len - 3 ; i++ {
		letter := bFace[i:i+1]
		if 4 == strings.Count(aFace, letter) {
			b = letter
			break
		}
	}
	return compareLetter(a,b)
}

//同花顺
func CompareStraightFlush(aFace,bFace string) (result int) {
	len := len(aFace)
	a := aFace[len - 1:len]
	b := bFace[len - 1:len]

	return compareLetter(a,b)
}

//通过牌面face比较大小
func compareByFace(aliceFace, bobFace string, pokerType int) (result int) {

	switch pokerType {
	case model.ROYALFLUSH:
		return model.EQUAL
	case model.STRAIGHTFLUSH:
		fmt.Println(aliceFace, bobFace)
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
		return
	case model.ONEPAIR:
		return
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

	// 解析牌，排序
	ClassifyCard(Alice)
	ClassifyCard(Bob)

	result = compareByType(Alice.Type, Bob.Type)
	if result == model.EQUAL {
		//同一类型牌面
		return compareByFace(Alice.SortFace, Bob.SortFace, Alice.Type)
	}

	return result
}
