package main

import "fmt"

func products(a uint32, b uint32) uint32 {
	return (a * b)
}

func multi() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = products(a, b)
	fmt.Println(res)
}
