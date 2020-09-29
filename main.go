package main

import "fmt"

func main()  {
	var base float32
	var altura float32
	fmt.Scan(&base)
	fmt.Scan(&altura)
	area := base * altura / 2
	fmt.Println(area)
}
