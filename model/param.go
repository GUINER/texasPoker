package model


//一手牌的牌个数
const CardAmount = 5
//用来表示一组卡牌所需的字符数
const AmountOfCardLetter = CardAmount * 2

//牌面的组成成员
const CardStr  = "23456789TJQKA"

//为牌面的符号编号
var CardLetters = map[string]int {"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"T":10,"J":11,"Q":12,"K":13,"A":14}

//牌的种类
var CardColor = []string{"s", "h", "c","d"}
var FaceList = []string{"2","3","4","5","6","7","8","9","T","J","Q","K","A"}

// 牌的种类
const (
	SPADE="s"	//黑桃
	HAERT="h"	//红心
	CLUB="c"	//草花
	DIAMOND="d"	//砖石
)



//牌面类型编号
const (
	ROYALFLUSH 	  	= 10		//皇家同花顺
	STRAIGHTFLUSH 	= 9			//同花顺
	FOUROFAKIND 	= 8			//四条
	FULLHOUSE 		= 7			//俘虏
	FlUSH			= 6			//同花
	STRAIGHT  		= 5			//顺子
	THREEOFAKIND 	= 4			//三条
	TWOPAIR 		= 3			//二对
	ONEPAIR 		= 2			//一对
	NOPAIR  		= 1			//单张
)

//牌面类型
var HandCardType = map[int]string{
	ROYALFLUSH		: "ROYAL FLUSH",	//皇家同花顺
	STRAIGHTFLUSH 	: "STRAIGHT FLUSH",	//同花顺
	FOUROFAKIND 	: "FOUR OF A KIND",	//四条
	FULLHOUSE 		: "FULL HOUSE",		//俘虏
	FlUSH			: "FlUSH",			//同花
	STRAIGHT  		: "STRAIGHT",		//顺子
	THREEOFAKIND 	: "THREE OF A KIND",//三条
	TWOPAIR 		: "TWO PAIR",		//二对
	ONEPAIR 		: "ONE PAIR",		//一对
	NOPAIR  		: "NO PAIR",		//单张
}
// 单张牌的信息
type CardFace struct {
	Face 	string 	`json:"face" fname:"牌面"`
	Color 	string 	`json:"color" fname:"牌色"`
	Index 	int 	`json:"index" fname:"牌的符号编号"`
}

// 一手扑克牌
type HandCards struct {
	Src 	string `json:"src" fname:"原始的一手牌"`
	Sort 	string `json:"sort" fname:"排序后的一手牌"`
	SortFace 	string `json:"sort_face" fname:"排序后的一手牌面"`
	SortColor 	string `json:"sort_color" fname:"排序后的一手牌色"`
	Type 	int	   `json:"type" fname:"牌面类型"`
}


// 比较结果
const (
	EQUAL = 0
	GREAT = 1
	LESS  = 2
)

type OnePoker struct {
	Alice 	string 	`json:"alice" fname:"Alice的牌"`
	Bob   	string 	`json:"bob" fname:"Bob的牌"`
	Result 	int 	`json:"result" fname："比较结果"`
}

type PokerData struct {
	Matchs []OnePoker `json:"matches" fname:"poker数据"`
}


