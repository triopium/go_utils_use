package main

import (
	"fmt"

	"github.com/triopium/go_utils/pkg/helper"
)

func main() {
	cdate := helper.TimeCurrent()
	fmt.Println(cdate)
	helper.Helloer()
}
