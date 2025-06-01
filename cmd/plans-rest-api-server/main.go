package main

import (
	"fmt"
	"plans-rest-api-server/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
