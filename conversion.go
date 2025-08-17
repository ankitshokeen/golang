package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Printf("rate our in between 1 to 10 \n")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating our pizza: ", input)

	// adding num value to input var which is byDefault init by string, it cannot be done because input is init as string and we are adding num value to it
	// inRating := input + 1

	// here it will be work fine we are doing type conversion
	inRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inRating + 1)
	}
}
