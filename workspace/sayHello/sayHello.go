package main

import ("fmt"
	"golang.org/x/example/hello/reverse")

func main(){
	fmt.Println("ORIGINAL : Hello, 34528")
	fmt.Println(reverse.String("Hello"), reverse.Int(34528))
}
