package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var params = os.Args
	if len(params) != 4 {
		fmt.Println("Argument count should be 3!")
	}

	param_a, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Error occured", params[1])
	}

	param_b, err := strconv.Atoi(params[3])
	if err != nil {
		fmt.Println("Error occured", params[3])
	}

	opcode := params[2]
	switch opcode {
		case "add", "plus":
		fmt.Println(param_a + param_b)
		case "ext", "minus":
		fmt.Println(param_a - param_b)
		case "mul":
		fmt.Println(param_a * param_b)
		case "div":
		fmt.Println(param_a / param_b)
		default:
		fmt.Println("unknown operator found. stopped, plz use add(+)/ext(-)/mul(*)/div(/)")
	}
}