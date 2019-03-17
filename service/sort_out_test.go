package service

import (
	"fmt"
	"texasPoker/model"
	"testing"
)

func TestSortTwoHandCard(t *testing.T) {
	alice := model.HandCards{Src:""}
	bob := model.HandCards{Src:""}

	if err := SortTwoHandCard(&alice, &bob); err != nil {
		fmt.Errorf("Error : %v", err)
	}
	fmt.Printf("alice[%s] - bob[%s]\n", alice.Sort, bob.Sort)
}
