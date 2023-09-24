package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	stringNum, _ := reader.ReadString('\n')
	floatNum, _ := strconv.ParseFloat(strings.TrimSpace(stringNum), 64)

	fmt.Printf("The truncated integer is: %f\n", math.Trunc(floatNum))
}
