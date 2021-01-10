package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"gostudy/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, contents map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

const url = "http://www.imooc.com"

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":   "robin",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}
func main() {
	r := &mock.Retriever{}
	fmt.Println(post(r))
	fmt.Println(session(r))
}
