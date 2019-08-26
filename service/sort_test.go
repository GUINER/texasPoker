package service

import (
	"fmt"
	"testing"
	"texasPoker/model"
	"time"
)

func TestSortTwoHandCard(t *testing.T) {
	time.Sleep(1 * time.Second)
	alice := model.HandCards{Src: "QdQd8c4s8c"}
	bob := model.HandCards{Src: "Qc9sAcAsQh"}

	timer("", func() { SortCard(&alice) })
	timer("", func() { SortCard(&bob) })

	fmt.Printf("alice[%s] - bob[%s]\n", alice.Sort, bob.Sort)
}

func TestString(t *testing.T) {
	var a []int
	var c []int
	var s []int
	timer("test", func() {
		for i := 0; i < 5; i++ {
			s = append(s, 2)
			a = append(a, 2)
			c = append(c, 2)
		}
	})
}
