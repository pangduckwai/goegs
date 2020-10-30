package main

import "fmt"

func test1(x interface{}) string {
	return fmt.Sprintf("%v", x)
}

func test2(y ...interface{}) string {
	return fmt.Sprintf("%v", y)
}

func main() {
	fmt.Println(test1(4))
	fmt.Println(test1(99))
	fmt.Println(test1(12))

	val := []uint{7, 8, 9}
	fmt.Println(test2(val))
}
