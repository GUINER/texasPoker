package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myproject/texasPoker/model"
	"myproject/texasPoker/service"
	"time"
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

func main() {
	beginTime := time.Now()

	Pokers := model.PokerData{}

	//1. 获取牌源
	GetPokerJsonData("texasPoker/data/match.json", &Pokers)


	for _,v := range Pokers.Matchs {
		var alice = model.HandCards{Src: v.Alice}
		var bob = model.HandCards{Src: v.Bob}

		//2. 两手牌进行比较
		//result := service.CompareTwoHandCard(&alice, &bob)
		service.CompareTwoHandCard(&alice, &bob)

		//3. 输出结果
		//fmt.Printf("%d. ",k)
		//service.OutputResult(alice.Src, bob.Src, result)

	}

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}