package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
        "strconv"
)

func main() {
	var grade float64
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	grade, err = strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)
}
