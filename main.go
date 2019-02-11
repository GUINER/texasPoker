package main

/*
*  德州扑克比较牌面大小
*  规则
	1. 牌号:2-10,J,Q,K,A,其中10为T,大小关系:从左往右递增
	2. 花色:花色黑桃spades，红心hearts，方块diamonds，草花clubs,即S(spades)、H(hearts)、D(diamonds)、C(clubs)
	牌面表示:FACE+color,例如: QsQhQdQcJh

	思路: 
		1.分级归类,根据牌面的类型,将多手牌分出级别,如:皇家同花顺>同花顺>四条>福满堂(三带二)>同花>顺子>三条>两队>一对>单张;
		2.实现每种类型的的牌比较大小的方法,并进行排序;

*
*/

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
		fmt.Errorf("ioutil.ReadFile %s, error: %v", filename, err)
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	//fmt.Print(v)
	if err != nil {
		fmt.Errorf("json.Unmarshal %s, error: %v", filename, err)
		return
	}
}

func main() {
	beginTime := time.Now()

	Pokers := model.PokerData{}

	//1. 获取牌源
	GetPokerJsonData("texasPoker/data/match.json", &Pokers)


	for k,v := range Pokers.Matchs {
		var alice = model.HandCards{Src: v.Alice}
		var bob = model.HandCards{Src: v.Bob}

		//2. 两手牌进行比较
		result := service.CompareTwoHandCard(&alice, &bob)

		//3. 输出结果
		fmt.Printf("%d. ",k)
		service.OutputResult(alice.Src, bob.Src, result)

	}

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}