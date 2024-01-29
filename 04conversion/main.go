package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza rating app.")
	fmt.Println("Please rate our pizza between 1 and 5.")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating - ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic(err) // Completely ends the whole program
	} else {
		fmt.Println("Added 1 to your rating - ", numRating+1)
	}

	fmt.Println("END OF PROGRAM") // wont be printed in case of error above cause of 'Panic'
}
