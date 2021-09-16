package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	mins, _ := strconv.Atoi(os.Args[1])

	price, err := GetPrice(mins)

	fmt.Printf("Minutes: %v\n", mins)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	} else {
		fmt.Printf("Price: %v\n", price)
	}
}
