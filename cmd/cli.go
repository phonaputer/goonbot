package main

import (
	"fmt"
	"goonbot/internal/rtd"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")

	fmt.Println(rtd.RollTheDice(input))
}
