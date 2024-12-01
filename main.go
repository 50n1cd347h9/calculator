package main
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var expr_str string
var op []string
var f int = 0

func main() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		f = 0
		expr_str = ""

		expr_str, _ = reader.ReadString('\n')

		if (len(expr_str) == 0) {
			break
		}

		var re = regexp.MustCompile(`[+\-*\/()]`)
		expr_str = re.ReplaceAllString(expr_str, " $0 ") // insert space
		op = strings.Fields(expr_str) // split by whitespace	
	
		var e = expr()
		fmt.Println("\t=", e)
	}
}

func expr() float64 {
	e := term()

	if (f >= len(op)) {
		return e
	}

	for f < len(op) && (op[f] == "+" || op[f] == "-") {
		if (op[f] == "+") {
			f++
			e =  e + term()
		} else {
			f++
			e = e - term()
		}
	}
	return e
}
func term() float64 {
	e := factor()

	for f < len(op) && (op[f] == "*" || op[f] == "/") {
		if (op[f] == "*") {
			f++
			e = e * factor()
		} else {
			f++
			e = e / factor()
		}
	}
	return e
}

func factor() float64 {
	var e float64

    pattern := `^[+-]?([0-9]+[.]?[0-9]*|[.][0-9]+)$`
    re, _ := regexp.Compile(pattern)
    match := re.MatchString(op[f])

	if (match) {
		var tmp = op[f]
		f++
		res, _ := strconv.ParseFloat(tmp, 64)
		return res
	} else if (op[f] == "(") {
		f++
		e = expr()
		if (op[f] != ")") {
			f++
			fmt.Println("error: missing ) at ", op[f], ":end")
		}
		return e
	} else {
		fmt.Println("error: expected number or ( at ", op[f], ":end")
		return 0
	}
}
