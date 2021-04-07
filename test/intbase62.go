package main

import (
	"fmt"
	"shorturl/util"
)

func main() {
	num := util.NewIntBase62("abcd")
	num.Acc()
	fmt.Println(num.String())
}
