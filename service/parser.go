package service

import (
	"myproject/texasPoker/model"
)


func ParseTwoHandCard(alice, bob *model.HandCards) {
	// 是否带有赖子
	alice.IsGhost, _ = HasGhost(alice.SortFace)
	bob.IsGhost, _ 	 = HasGhost(bob.SortFace)

	//1. 5张
	if model.FiveCard == len(alice.SortFace) {
		return
	}
	//2. 7张
	if model.SevenCard == len(alice.SortFace) {
		SevenCardParse(alice)
		SevenCardParse(bob)
		return
	}

	return
}