package service

import (
	"fmt"
	"testing"
)

func TestCardIsStraight(t *testing.T){
	if isOk,err := CardIsStraight("23456"); err != nil {
		fmt.Printf("非顺子\n")
	} else {
		if isOk {
			fmt.Printf("顺子\n")
		}
	}
}


