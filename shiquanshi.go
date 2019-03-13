package main

/*
	题目: 2013年,月日时分秒组成的10位数,其中是十全时的数的有多少个?例如:0928175643
	条件:
		1. 十全时数的数字不能重复,0-9各占一个数;
		2. 在2003年的时间中筛选;
	过滤:
		1. 月份1~12,天数1~31,时:01~23,分:01~59,秒:01~59;
		2. 重复出现的数字,如00,11,22...类似的数字;
*/
import (
	"fmt"
	"strings"
	"strconv"
)

/*
*功能: 判断是否为十全时
*param: shiquan
*return: bool
*/
func IsShiQuanShi(shiquan string) bool {
	for i := 0; i < 10 ; i++ {
		a := strconv.Itoa(i)
		count := strings.Count(shiquan, a)	//该数字在shiquan出现的次数
		//fmt.Printf("%s,%d,%d\n", shiquan, i, count)
		if  count > 1 {
			return false
		}
	}
	return true
}


func main() {

	CountShiQuanShi := 0

	for month := 1; month <= 12; month ++  {	//月份范围
		//fmt.Printf("month:%d\n", month)
		if month == 11 {	//11过滤掉
			continue
		}
		if month <= 12 {
			for day := 1; day <= 31; day ++ {		//天数范围
				if ( day / 11 != 0 ) && ( day % 11 == 0 ) { //11,22的数字可以过滤掉
					continue
				}
				if ( month == 2 ) && day > 28 { //2月
					continue
				} else if ( month == 4 || month == 6 || month == 9 || month == 11 ) && day > 30 { // 30天
					continue
				//} else if ( month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 ) && day > 31 {	//31天
				//	continue
				}

				for hour := 1; hour < 24 ; hour ++ {	//小时范围
					if ( hour / 11 != 0 ) && ( hour % 11 == 0 ) {
						continue
					}
					for min := 1; min < 60 ; min ++ {	//分钟范围
						if ( min / 11 != 0 ) && ( min % 11 == 0 ) {
							continue
						}
						for second := 1; second < 60 ; second ++ {	//秒钟范围
							if ( second / 11 != 0 ) && ( second % 11 == 0 ) {
								continue
							}
							ShiQuan := fmt.Sprintf("%02d%02d%02d%02d%02d", month, day, hour, min, second)
							if true == IsShiQuanShi(ShiQuan) {	// 判断是否为十全时
								CountShiQuanShi++
								fmt.Printf("%d. %s\n", CountShiQuanShi, ShiQuan )
							}
						}
					}
				}

			}
		}

	}
	fmt.Printf("total shiquanshi: %d\n", CountShiQuanShi)

}
