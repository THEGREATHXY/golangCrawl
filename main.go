package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url: "https://book.douban.com",
		ParseFunc: parse.ParseTag,
		//Url: "https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4",
		//ParseFunc: parse.ParseBookList,
		//Url: "https://book.douban.com/subject/10554308/",
		//ParseFunc: parse.ParseBookDetail,
	})

}


