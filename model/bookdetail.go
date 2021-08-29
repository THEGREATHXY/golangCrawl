package model

type BookDetail struct {
	BookName string
	Author string
	Publisher string
	BookPages string
	Price string
	Score string
	Introduce string
}

func (b BookDetail) String() string {
	return "书名：" + b.BookName + " 作者：" + b.Author + " 出版社：" + b.Publisher +
		" 页数：" + b.BookPages + " 价格：" + b.Price + " 得分：" + b.Score +
		"\n简介：" + b.Introduce
}