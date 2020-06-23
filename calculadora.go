package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Calc struct {
}

func GetOperator(in string) string {
	var validID = regexp.MustCompile(`[\+*-/]`)
	fmt.Println(validID.FindString(in))
	return validID.FindString(in)
}

func (Calc) operate(entrada string, operator string) int {
	clean_input := strings.Split(entrada, operator)
	operator1 := parsear(clean_input[0])
	operator2 := parsear(clean_input[1])
	switch operator {
	case "-":
		return operator1 - operator2
	case "+":
		return operator1 + operator2
	case "*":
		return operator1 * operator2
	case "/":
		return operator1 / operator2
	default:
		fmt.Println("operador no soportado")
		return 0
	}
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text_ := scanner.Text()
	return text_
}
func parsear(entrada string) int {
	input, _ := strconv.Atoi(entrada)
	return input
}
