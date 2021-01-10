package main

import "fmt"

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err.Error())
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	//panic("123")

	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
