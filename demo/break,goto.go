package main

import "fmt"

//func main() {
////	for i := 0; i < 5; i++ {
////		for j := 0; j < 10; j++ {
////			if j>5 {
////				goto label
////			}
////			fmt.Println(j)
////		}
////	}
////label:
////	fmt.Println("AAAAA")
////}
func main() {
label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j>5 {
				break label
			}
			fmt.Println(j)
		}
	}
	fmt.Println("AAAAA")
}
