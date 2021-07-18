package grug

import (
	"math"
	"strconv"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			// Computes the result of arg 0 + arg 1
			Name: "Plus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return a + b, nil
			},
		},
		{
			// Computes the result of arg 0 - arg 1
			Name: "Minus",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return a - b, nil
			},
		},
		{
			// Computes the result of arg 0 + arg 1
			Name: "Multiply",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return a * b, nil
			},
		},
		{
			// Computes the result of arg 0 + arg 1
			Name: "Divide",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return a / b, nil
			},
		},
		{
			// Computes the result of arg 0^arg 1
			Name: "Pow",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				b, _ := strconv.Atoi(atostr(args[1]))
				return math.Pow(float64(a), float64(b)), nil
			},
		},
		{
			// Increments the argument by 1
			Name: "Increment",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				return a + 1, nil
			},
		},
		{
			// Increments the argument by 1
			Name: "Decrement",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, _ := strconv.Atoi(atostr(args[0]))
				return a - 1, nil
			},
		},
	}...)
}
