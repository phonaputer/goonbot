package main

import (
	"fmt"
	"goonbot/internal/goonbot/rtd"
	"os"
)

func main() {
	fmt.Println(rtd.RollTheDice(os.Args[1:]))
}
