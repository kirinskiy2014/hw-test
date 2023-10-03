package main

import (
	"fmt"

	"github.com/kirinskiy2014/hw-test/hw01_hello_otus/stringutil"
)

func main() {
	originString := "Hello, OTUS!"
	reversedString := stringutil.Reverse(originString)
	fmt.Println(reversedString)
}
