package main

import (
	"fmt"
	"myproject/texasPoker/service"
	"time"
)



func main() {
	var beginTime time.Time
	var finishTime time.Time

	fmt.Print("5张无赖子 ")
	beginTime = time.Now()
	service.StartService("texasPoker/data/match.json")
	finishTime = time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)


	fmt.Print("5张有赖子 ")
	beginTime = time.Now()
	service.StartService("texasPoker/data/five_cards_with_ghost.json")
	finishTime = time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)

	fmt.Print("7张无赖子 ")
	beginTime = time.Now()
	service.StartService("texasPoker/data/seven_cards.json")
	finishTime = time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)


	fmt.Print("7张有赖子 ")
	beginTime = time.Now()
	service.StartService("texasPoker/data/seven_cards_with_ghost.json")
	finishTime = time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}