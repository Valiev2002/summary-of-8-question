package main

import "fmt"

func main() {
	fibonacci(4)
}
func myPrint(a ...interface{}) {
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}

}
