package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	// greetName()
	// stringParsing()
	// stringMethods()
	dateMethods()
}

func dateMethods() {
	date := time.Now()
	fmt.Printf("Date: %s %d, %d\n", date.Month().String(), date.Day(), date.Year())
	fmt.Printf("Time: %d:%d:%d\n", date.Hour(), date.Minute(), date.Second())
	fmt.Println(date.Weekday())
}

func stringMethods() {
	word := "onomatopoeia"

	// Does not modify original value
	// Case sensitive
	fmt.Println(strings.ReplaceAll(word, "o", "0"))
	fmt.Println(strings.Contains(word, "0"))
	fmt.Println(strings.Index(word, "o")) //returns first instance
	fmt.Println(strings.Split(word, "o")) // returns []string
	fmt.Println(strings.ToUpper(word))
	fmt.Println(strings.HasSuffix(word, "poeia"))

	// Print each unicode character
	for i, rune := range word {
		fmt.Printf("%d : %#U\n", i, rune)
	}
}

func stringParsing() {
	// strconv package has required methods
	num := "10000"
	numAsInt, _ := strconv.Atoi(num)
	fmt.Println(num, reflect.TypeOf(num))
	fmt.Println(numAsInt, reflect.TypeOf(numAsInt))

	decimal := "3.14259"
	decAsFloat, err := strconv.ParseFloat(decimal, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decAsFloat, reflect.TypeOf(decAsFloat))

	fmt.Println(strings.Compare(num, decimal) == 0)
	fmt.Println(strings.Compare(num, "10000") == 0)

}

func greetName() {
	fmt.Print("Hello! Please enter you name: ")
	name := captureInput()
	for range name {
		fmt.Printf("Hello, %s! \n", name)
	}
}

func captureInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	// Remove newline at the end of reader
	return strings.TrimSpace(input)
}
