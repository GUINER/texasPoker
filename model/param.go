package model


//一手牌的牌种类
const (
	FiveCard 	= 5
	FiveCardWithGhost =  6
	SevenCard 	= 7
	SevenCardWithGhost = 8
)

const CardAmount = FiveCard
//用来表示一组卡牌所需的字符数
const AmountOfCardLetter = CardAmount * 2

//牌面的组成成员
const CardStr  = "A23456789TJQKA"


//为牌面的符号编号
var CardLetters = map[string]int {"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"T":10,"J":11,"Q":12,"K":13,"A":14,"X":15}
var FaceList = []string{"2","3","4","5","6","7","8","9","T","J","Q","K","A","X"}
const (
	Ace 	= "A"
	Two 	= "2"
	Three 	= "3"
	Four 	= "4"
	Five 	= "5"
	Six 	= "6"
	Seven 	= "7"
	Eight 	= "8"
	Nine 	= "9"
	Ten 	= "T"
	Jazz 	= "J"
	Queen 	= "Q"
	King 	= "K"
	Ghost 	= "X"
)

//牌的花色
const (
	SPADE	=	"s"	//黑桃
	HEART	=	"h"	//红心
	CLUB	=	"c"	//草花
	DIAMOND	=	"d"	//砖石
	GhostColor	=	"n"	//赖子花色
)
var CardColor = []string{SPADE, HEART, CLUB, DIAMOND, GhostColor}



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
	UNKNOWN			= 0
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
	UNKNOWN			: "Unknown type",	//未知类型
}


//顺子
const (
	SA2345 = "2345A"
	S23456 = "23456"
	S34567 = "34567"
	S45678 = "45678"
	S56789 = "56789"
	S6789T = "6789T"
	S789TJ = "789TJ"
	S89TJQ = "89TJQ"
	S9TJQK = "9TJQK"
	STJQKA = "TJQKA"
)
// 顺子等级
var StraightList = map[string]int{
	SA2345 : 1,
	S23456 : 2,
	S34567 : 3,
	S45678 : 4,
	S56789 : 5,
	S6789T : 6,
	S789TJ : 7,
	S89TJQ : 8,
	S9TJQK : 9,
	STJQKA : 10,
}
//顺子编号
var StraightRankList = []string{
	SA2345, S23456, S34567 ,S45678 ,S56789 ,S6789T ,S789TJ ,S89TJQ, S9TJQK, STJQKA,
}


// 单张牌的信息
type CardFace struct {
	Face 	string 	`json:"face" fname:"牌面"`
	Color 	string 	`json:"color" fname:"牌色"`
	Index 	int 	`json:"index" fname:"牌的符号编号"`
}

// 一手扑克牌
type HandCards struct {
	Src 		string `json:"src" fname:"原始的一手牌"`
	Sort 		string `json:"sort" fname:"排序后的一手牌"`
	SortFace 	string `json:"sort_face" fname:"排序后的一手牌面"`
	SortColor 	string `json:"sort_color" fname:"排序后的一手牌色"`
	Type 		int	   `json:"type" fname:"牌面类型"`
	IsGhost		bool   `json:"is_ghost" fname:"是否有赖子"`
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


