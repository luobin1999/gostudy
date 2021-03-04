package main

import (
	"fmt"
	"gostudy/crawler/engine"
	"gostudy/crawler/zhenai/parser"
)

const url = "https://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
	//fun1()
}

//验证给切片中添加nil值，切片长度不变
func fun1() {
	request := []engine.Request{}
	fmt.Println(len(request))
	//直接给切片中添加nil值会编译失败，所以构造空切片进行添加
	result := engine.ParserResult{
		Request: nil,
		Items:   nil,
	}
	request = append(request, result.Request...)
	fmt.Println(len(request))
	/* 直接向切片中添加nil值会编译失败
	request = append(request, engine.Request(nil))
	fmt.Println(len(request))
	*/
}
