package calc

import (
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	var (
		numb                    string
		exp, stap               []string
		multdiv, staples, trash int
		endstap, begstap        []int
		tf                      bool
		err                     error
		a, f                    float64
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
			return 0, fmt.Errorf("empty string")
		} else {
			a, err := strconv.ParseFloat(string(exp[0]), 64)
			if err != nil {
				return 0, fmt.Errorf("strange expression")
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

	for staples != len(begstap) {
		if len(begstap) != 0 {
			stap = exp[begstap[staples]+1 : endstap[staples]]
		}

		for i := 0; i < len(stap); i++ {
			if stap[i] == "/" || stap[i] == "*" {
				multdiv++
			}
			tf, err = findErrors(exp, stap)

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
					multdiv--
					trash++

				} else if stap[i] == "*" {
					a, _ = strconv.ParseFloat(stap[i-1], 64)
					f, _ = strconv.ParseFloat(stap[i+1], 64)
					numb = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
					stap[i-1] = numb
					multdiv--
					trash++

				} else if stap[i] == "-" && multdiv == 0 {
					a, _ = strconv.ParseFloat(stap[i-1], 64)
					f, _ = strconv.ParseFloat(stap[i+1], 64)
					numb = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
					stap[i-1] = numb
					trash++

				} else if stap[i] == "+" && multdiv == 0 {
					a, _ = strconv.ParseFloat(stap[i-1], 64)
					f, _ = strconv.ParseFloat(stap[i+1], 64)
					numb = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
					stap[i-1] = numb
					trash++
				}

				if i >= 1 {
					if stap[i-1] == numb {
						copy(stap[i+1:], stap[i+2:])
						copy(stap[i:], stap[i+1:])
						stap = stap[:len(stap)-2]
						numb = ""
					}
				}

				if trash == 1 {
					i = 0
					trash = 0
				}
			}

			if len(stap) > 0 {
				copy(exp[begstap[staples]+2:], exp[endstap[staples]+1:])
				copy(exp[begstap[staples]:], exp[begstap[staples]+1:])
				exp = exp[:len(exp)-(endstap[staples]-begstap[staples])]
			}

			if exp[len(exp)-1] == " " || exp[len(exp)-1] == "" {
				exp = exp[:len(exp)-1]
			}
			staples++

			if staples < len(begstap) {
				if begstap[staples]-4 > 0 {
					for i := 0; i < len(begstap); i++ {
						endstap[i] -= 4
						begstap[i] -= 4
						if begstap[i] < 0 {
							begstap[i] = 0
						}
					}
				}
			}
		}
	}
	//ВЫРАЖЕНИЕ!!!!!!!!!!/////////////////////////////////////////
	tf, err = findErrors(exp, stap)
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
				multdiv--
				trash++

			} else if exp[i] == "*" {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
				exp[i-1] = numb
				multdiv--
				trash++

			} else if exp[i] == "-" && multdiv == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
				exp[i-1] = numb
				trash++

			} else if exp[i] == "+" && multdiv == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				numb = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
				exp[i-1] = numb
				trash++
			}

			if i >= 1 {
				if exp[i-1] == numb {
					copy(exp[i+1:], exp[i+2:])
					copy(exp[i:], exp[i+1:])
					exp = exp[:len(exp)-2]
				}
			}

			if trash == 1 {
				i = 0
				trash = 0
			}
		}
	}

	a, _ = strconv.ParseFloat(exp[0], 64)
	return a, nil
}

func findErrors(exp, stap []string) (bool, error) {

	for i := 0; i < len(exp); i++ {
		if i == 0 {
			if (exp[i] < "0" || exp[i] > "9") && exp[i] != "(" {
				return true, fmt.Errorf("error in writing, the sign at the beginning")
			}

		}

		if i == len(exp)-1 {
			if exp[i] == "/" || exp[i] == "*" || exp[i] == "+" || exp[i] == "-" || exp[i] == "(" {
				return true, fmt.Errorf("error in writing, the sign at the end")
			} else {
				if (exp[i] == ")") && (exp[i-1] == "*" || exp[i-1] == "(" || exp[i-1] == "+" || exp[i-1] == "-" || exp[i-1] == "/") {
					return true, fmt.Errorf("error in writing, sign before staple")
				} else {
					return false, fmt.Errorf("Good")
				}
			}
		}

		if i == 0 {
			if exp[i] == "(" {
				if exp[i+1] == "+" || exp[i+1] == "-" || exp[i+1] == "*" || exp[i+1] == "/" || exp[i+1] == ")" {
					return true, fmt.Errorf("error in writing, sign after staple")
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
			_, b := strconv.ParseFloat(exp[i-1], 64)

			if exp[i+1] == "(" || exp[i] == "(" {
				if exp[i+1] == "(" {
					if b != nil {
						return true, fmt.Errorf("error in writing, staples")
					}
				} else if exp[i] == "(" {
					if a != nil || b == nil {
						return true, fmt.Errorf("error in writing, staples")
					}
				}
				continue
			}

			if a != nil || b != nil {
				return true, fmt.Errorf("error in writing, incorrect symbol after/before sign")
			}

			if exp[i] == "/" {
				if exp[i+1] == "0" {
					return true, fmt.Errorf("divide by zero")
				}
			}
		}
	}
	return false, fmt.Errorf("Good")
}
