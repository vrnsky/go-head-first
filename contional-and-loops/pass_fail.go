package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
        "strings"
)

func main() {
	var grade float64
        var status string
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
        input = strings.TrimSpace(input)
	grade, err = strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
