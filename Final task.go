package main

import (
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	var (
		numb             string
		exp, stap        []string
		multdiv, bc      int
		endstap, begstap []int
		tf               bool
		err              error
		a, f             float64
	)

	for i := 0; i < len(expression); i++ {
		if string(expression[i]) != "/" && string(expression[i]) != "*" && string(expression[i]) != "+" && string(expression[i]) != "-" && string(expression[i]) != "(" && string(expression[i]) != ")" {
			numb += string(expression[i])
		} else {
			if numb != "" {
				exp = append(exp, numb)
				numb = ""
			}
			exp = append(exp, string(expression[i]))
		}
	}

	exp = append(exp, numb)

	if len(exp) <= 1 {
		if len(exp) == 0 {
			return 0, fmt.Errorf("Empty string")
		} else {
			a, err := strconv.ParseFloat(string(exp[0]), 64)
			if err != nil {
				return 0, fmt.Errorf("Error in writing")
			} else {
				return a, nil
			}
		}
	}

	/////////СКОБКИ!!!!!!!!////////////////////////////////////////////////////////////////////////////

	for i := 0; i < len(exp); i++ {
		if exp[i] == "(" {
			begstap = append(begstap, i)
		} else if exp[i] == ")" {
			endstap = append(endstap, i)
		}
	}

	for bc != len(begstap) {
		if len(begstap) != 0 {
			stap = exp[begstap[bc]+1 : endstap[bc]]
			fmt.Println(stap, begstap, endstap, exp)
		}

		for i := 0; i < len(stap); i++ {
			if stap[i] == "/" || stap[i] == "*" {
				multdiv++
			}
		}

		tf, err = FindErrors(begstap[bc], endstap[bc], exp, stap)
		if tf {
			return 0, err
		}
	}
	for len(stap) > 1 {

		if len(stap) == 0 {
			break
		}

		for i := 0; i < len(stap); i++ {
			if stap[i] == "/" {
				a, _ = strconv.ParseFloat(stap[i-1], 64)
				f, _ = strconv.ParseFloat(stap[i+1], 64)
				numb = strconv.FormatFloat(float64(a)/float64(f), 'f', 3, 64)
				stap[i-1] = numb
				copy(stap[i+1:], stap[i+2:])
				copy(stap[i:], stap[i+1:])
				stap = stap[:len(stap)-2]
				multdiv--
				i = 0

			} else if stap[i] == "*" {
				a, _ = strconv.ParseFloat(stap[i-1], 64)
				f, _ = strconv.ParseFloat(stap[i+1], 64)
				numb = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
				stap[i-1] = numb
				copy(stap[i+1:], stap[i+2:])
				copy(stap[i:], stap[i+1:])
				stap = stap[:len(stap)-2]
				multdiv--
				i = 0

			} else if stap[i] == "-" && multdiv == 0 {
				a, _ = strconv.ParseFloat(stap[i-1], 64)
				f, _ = strconv.ParseFloat(stap[i+1], 64)
				numb = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
				stap[i-1] = numb
				copy(stap[i+1:], stap[i+2:])
				copy(stap[i:], stap[i+1:])
				stap = stap[:len(stap)-2]
				i = 0

			} else if stap[i] == "+" && multdiv == 0 {
				a, _ = strconv.ParseFloat(stap[i-1], 64)
				f, _ = strconv.ParseFloat(stap[i+1], 64)
				numb = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
				stap[i-1] = numb
				copy(stap[i+1:], stap[i+2:])
				copy(stap[i:], stap[i+1:])
				stap = stap[:len(stap)-2]
				i = 0
			}
		}
		if len(stap) > 0 {
			copy(exp[begstap[bc]+2:], exp[endstap[bc]+1:])
			copy(exp[begstap[bc]:], exp[begstap[bc]+1:])
			exp = exp[:len(exp)-(endstap[bc]-begstap[bc])]
		}

		if exp[len(exp)-1] == " " || exp[len(exp)-1] == "" {
			exp = exp[:len(exp)-1]
		}
		bc++
		if bc < len(begstap) {
			if begstap[bc]-4 > 0 {
				endstap[bc] = endstap[bc] - 4
				begstap[bc] = begstap[bc] - 4
			}
		}
	}

	//ВЫРАЖЕНИЕ!!!!!!!!!!/////////////////////////////////////////

	tf, err = FindErrors(0, 0, exp, stap)
	if tf {
		return 0, err
	}

	for i := 0; i < len(exp); i++ {
		if exp[i] == "*" || exp[i] == "/" {
			multdiv++
		}
	}
	for len(exp) > 1 {
		for i := 0; i < len(exp); i++ {
			if exp[i] == "/" {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)/float64(f), 'f', 3, 64)
				exp[i-1] = numb
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				multdiv--
				i = 0
			} else if exp[i] == "*" {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
				exp[i-1] = numb
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				multdiv--
				i = 0
			} else if exp[i] == "-" && multdiv == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
				exp[i-1] = numb
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				i = 0
			} else if exp[i] == "+" && multdiv == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
				exp[i-1] = numb
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				i = 0
			}
		}
	}
	a, _ = strconv.ParseFloat(exp[0], 64)
	return a, nil
}
func main() {
	var expression string
	fmt.Scanln(&expression)
	fmt.Println(Calc(expression))
}
func FindErrors(begstap, endstap int, exp, stap []string) (bool, error) {
	for i := 0; i < len(exp); i++ {
		if i == len(exp)-1 {
			if exp[i] == "/" || exp[i] == "*" || exp[i] == "+" || exp[i] == "-" || exp[i] == "(" {
				return true, fmt.Errorf("Error in writing")
			} else {
				if (exp[i] == ")") && (exp[i-1] == "*" || exp[i-1] == "(" || exp[i-1] == "+" || exp[i-1] == "-" || exp[i-1] == "/") {
					return true, fmt.Errorf("Error in writing znack")
				} else {
					return false, fmt.Errorf("Good")
				}
			}
		}
		if i == 0 {
			if exp[i] == "(" {
				if exp[i+1] == "+" || exp[i+1] == "-" || exp[i+1] == "*" || exp[i+1] == "/" || exp[i+1] == ")" {
					return true, fmt.Errorf("Error in writing zn after scobs")
				} else {
					i++
				}
			}
		}
		if exp[i] == "/" || exp[i] == "*" || exp[i] == "+" || exp[i] == "-" {
			if exp[i+1] == "(" || exp[i-1] == ")" {
				continue
			}
			_, a := strconv.ParseFloat(exp[i+1], 64)
			_, endstap := strconv.ParseFloat(exp[i-1], 64)
			if exp[i+1] == "(" || exp[i] == "(" {
				if exp[i+1] == "(" {
					if endstap != nil {
						return true, fmt.Errorf("Error in writing scobs")
					}
				} else if exp[i] == "(" {
					if a != nil || endstap == nil {
						return true, fmt.Errorf("Error in writing scobs")
					}
				}
				continue
			}
			if a != nil || endstap != nil {
				return true, fmt.Errorf("Error in writing")
			}
			if exp[i] == "/" {
				if exp[i+1] == "0" {
					return true, fmt.Errorf("Divide by zero")
				}
			}
		}
	}
	return false, fmt.Errorf("Good")
}
