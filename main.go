package main

import (
	"fmt"

	"github.com/Luis-Designs/hello/greet"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.Spanish())
	fmt.Println(quote.HelloV3())
	fmt.Println(quote.Concurrency())
}
