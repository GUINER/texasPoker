package service

import (
	"fmt"
	"testing"
	"texasPoker/model"
)

func TestSortTwoHandCard(t *testing.T) {
	alice := model.HandCards{Src: "QdQd8c4s8c"}
	bob := model.HandCards{Src: "Qc9sAcAsQh"}

	timer("", func() { SortCard(&alice) })
	timer("", func() { SortCard(&bob) })

	fmt.Printf("alice[%s] - bob[%s]\n", alice.Sort, bob.Sort)
}
