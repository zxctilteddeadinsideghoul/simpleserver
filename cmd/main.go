package main

import (
	"fmt"
	"simpleserver/internal/config"
)

func main() {
	//config init
	cfg := config.MustLoad()

	fmt.Println(cfg)

}
