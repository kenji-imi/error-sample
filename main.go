package main

import (
	"fmt"

	kjerror "error-sample/error"
)

func main() {
	fmt.Println("[Error]")
	fmt.Printf("  - %+v\n", kjerror.Wrap1())
}
