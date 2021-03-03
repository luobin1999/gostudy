package main

import (
	"fmt"
	"regexp"
)

const text = "My email is luobin1999@foxmail.com"

const text1 = `My email is luobin1999@foxmail.com
email1 is abcd@qq.com
   email2 is 124@gmail.com   987ut@ahcv.com.cn`

const text2 = `My email is luobin1999@foxmail.com
email1 is abcd@qq.com
   email2 is 124@gmail.com   987ut@ahcv.com.cn`

/**
正则表达式
*/
func main() {
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	//匹配单个邮箱
	match := compile.FindString(text)
	//匹配所有邮箱
	match2 := compile.FindAllString(text1, -1)
	fmt.Println(match)
	fmt.Println(match2)

	//匹配所有邮箱并且将各个部分信息挖出来（加括号）
	compile1 := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match3 := compile1.FindAllStringSubmatch(text2, -1)
	for _, m := range match3 {
		fmt.Println(m)
	}
}
