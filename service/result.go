package service

import (
	"fmt"
	"myproject/texasPoker/model"
)

func OutputResult(alice,bob string, result int){
	if result == model.EQUAL {
		fmt.Printf("result: alice[%s] = bob[%s]\n", alice, bob)
	} else if result == model.GREAT {
		fmt.Printf("result: alice[%s] > bob[%s]\n", alice, bob)
	} else if result == model.LESS {
		fmt.Printf("result: alice[%s] < bob[%s]\n", alice, bob)
	} else {
		fmt.Printf("unknown result[%d]: alice[%s] - bob[%s]\n", result, alice, bob)
	}
}
