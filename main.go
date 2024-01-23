package main

import (
	"BagasA11/go-POS/configs"
	"fmt"
)

func main() {
	fmt.Println("DB init: ", configs.InitDb())
	fmt.Println("dsn: ", configs.Dsn)
	fmt.Println("DB: ", configs.GetDB())
}
