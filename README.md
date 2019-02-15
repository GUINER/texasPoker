# texasPoker


*  德州扑克比较牌面大小
*  规则
	1. 牌号:2-10,J,Q,K,A,其中10为T,大小关系:从左往右递增
	2. 花色:花色黑桃spades，红心hearts，方块diamonds，草花clubs,即S(spades)、H(hearts)、D(diamonds)、C(clubs)
	牌面表示:FACE+color,例如: QsQhQdQcJh

*   思路: 
	
		1.分级归类,根据牌面的类型,将多手牌分出级别,如:
		    
		    皇家同花顺 > 同花顺 > 四条 > 福满堂(三带二) > 同花 > 顺子 > 三条 > 两队 > 一对 > 单张;
		    
		2.实现每种类型的的牌比较大小的方法,并进行排序;
		
		3.赖子怎么分类？
		    皇家同花顺： TJQKAX, TJQKX, TJQAX, TJKAX, TQKAX, JQKAX
