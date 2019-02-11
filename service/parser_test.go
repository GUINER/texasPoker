package service

import (
	"fmt"
	"testing"
)

func TestSevenCardHasStraight(t *testing.T) {
	sevenCardFace := "2333456"
	isOk := SevenCardHasStraight(sevenCardFace)
	if isOk {
		fmt.Printf("[%s] 是顺子", sevenCardFace)
	} else {
		fmt.Printf("[%s] 不是顺子", sevenCardFace)
	}
}
