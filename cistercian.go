package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("Failed to parse input as an integer:", os.Args[1])
		os.Exit(1)
	}

	renderer := TextRenderer{Value: int(value)}
	renderer.Output()
}
