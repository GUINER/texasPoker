package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"texasPoker/model"
	"time"
)

func GetPokerJsonData(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ioutil.ReadFile %s, error: %v\n", filename, err)
		return

	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Printf("json.Unmarshal %s, error: %v\n", filename, err)
		return
	}
}

func timer(name string, f func()) {
	fmt.Print(name)

	beginTime := time.Now()

	f()

	finishTime := time.Now()
	fmt.Printf("共耗时：%.6f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}

func Timer(name string, file string) {
	Pokers := model.PokerData{}
	GetPokerJsonData(file, &Pokers)

	fmt.Print(name)

	beginTime := time.Now()

	// start poker
	StartPoker(&Pokers)

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}

func StartPoker(Pokers *model.PokerData) {
	for _, v := range Pokers.Matchs {
		comparer := NewPokerComparer(v.Alice, v.Bob)

		comparer.Compare()

		//fmt.Printf("%d. ",k)
		//comparer.PrintResult()
	}
}
