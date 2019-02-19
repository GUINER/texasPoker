package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myproject/texasPoker/model"
)

func GetPokerJsonData(filename string, v interface{}) {
	//JsonParse := model.JsonStruct{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ioutil.ReadFile %s, error: %v\n", filename, err)
		return

	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	//fmt.Print(v)
	if err != nil {
		fmt.Printf("json.Unmarshal %s, error: %v\n", filename, err)
		return
	}
}



func StartService(filename string) {
	Pokers := model.PokerData{}

	//1. 获取牌源
	GetPokerJsonData(filename, &Pokers)

	for _,v := range Pokers.Matchs {
		var alice = model.HandCards{Src: v.Alice}
		var bob = model.HandCards{Src: v.Bob}

		//2. 两手牌进行比较
		result := CompareTwoHandCard(&alice, &bob)
		//CompareTwoHandCard(&alice, &bob)

		//3. 输出结果
		//fmt.Printf("%d. ",k)
		//OutputResult(alice.Src, bob.Src, result)
		if -1 == result {
			fmt.Println(alice)
			fmt.Println(bob)
			fmt.Println()
		}

	}
}
