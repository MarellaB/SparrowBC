package main

import (
	"blockchain-golang/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.FindNeighbors("127.0.0.1", 5001, 0, 3, 5001, 5003))
}
