package main

import (
	"texasPoker/service"
)



func main() {

	service.Timer("5张无赖子 ", "data/match.json")

	service.Timer("5张有赖子 ", "data/five_cards_with_ghost.json")

	service.Timer("7张无赖子 ", "data/seven_cards.json")

	service.Timer("7张有赖子 ", "data/seven_cards_with_ghost.json")

}