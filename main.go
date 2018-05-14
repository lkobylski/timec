package main

import "fmt"
import (
	"os"
	"time"
	"strconv"
)

//args
var (
	timeIn int64
)

func main() {

	if len(os.Args) < 2{
		printUsage()
	}

	timeIn, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		printUsage()
	}

	t := time.Unix(timeIn, 0)
	fmt.Println(t)

	os.Exit(1)

}


func printUsage() {
	fmt.Printf("Usage: %s [timestamp]\n", os.Args[0])
	fmt.Print("timestamp must be integer\n", )
	os.Exit(1)
}