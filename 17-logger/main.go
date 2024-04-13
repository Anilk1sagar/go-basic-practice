package main

import (
	"fmt"

	logger "gitlab.com/anilk1sagar/go_basic_practice/17-logger/logr"
)

func main() {

	fmt.Println("welicome")

	logger.Log().Info("hello")

}
