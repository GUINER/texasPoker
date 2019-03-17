package service

import (
	"fmt"
	"texasPoker/model"
)

func OutputResult(alice,bob *model.HandCards, result int){
	if result == model.EQUAL {
		fmt.Printf("result: alice[%s][%s] = bob[%s][%s]\n", alice.Src, model.HandCardType[alice.Type], bob.Src, model.HandCardType[bob.Type])
	} else if result == model.GREAT {
		fmt.Printf("result: alice[%s][%s] > bob[%s][%s]\n", alice.Src, model.HandCardType[alice.Type], bob.Src, model.HandCardType[bob.Type])
	} else if result == model.LESS {
		fmt.Printf("result: alice[%s][%s] < bob[%s][%s]\n", alice.Src, model.HandCardType[alice.Type], bob.Src, model.HandCardType[bob.Type])
	} else {
		fmt.Printf("unknown result[%d]: alice[%s][%s] - bob[%s][%s]\n", result, alice.Src, model.HandCardType[alice.Type], bob.Src, model.HandCardType[bob.Type])
	}
}
