package service

import (
	"fmt"
	//"myproject/texasPoker/service"
	"testing"
)

func TestCompareStraightFlush(t *testing.T) {
	card1 := "AsKsQsJsTs"
	card2 := "AsKsQsJsTs"
	result := CompareStraightFlush(card1, card2)

	fmt.Printf("result: %d\n",result)
}
