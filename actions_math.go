package grug

import (
	"fmt"
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
		{
			// Gets the "maximum object" by checking against a given field
			Name: "MaxByField",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				objSlice, ok := args[0].([]interface{})
				if !ok {
					return nil, fmt.Errorf("unable to interpret %v as a slice", args[0])
				}

				accessPattern, ok := args[1].(string)
				if !ok {
					return nil, fmt.Errorf("unable to interpret %v as a string", args[1])
				}

				var maxObj interface{}
				maxVal := -1
				for _, o := range objSlice {
					field, err := fieldAccess(o, accessPattern)
					if err != nil {
						return nil, err
					}

					var val int
					switch t := field.(type) {
					case float32:
						val = int(t)
					case float64:
						val = int(t)
					case int:
						val = t
					default:
						return nil, fmt.Errorf("unable to interpret %v as a number", args[0])
					}

					if val > maxVal {
						maxVal = val
						maxObj = o
					}
				}

				if maxObj == nil {
					return nil, fmt.Errorf("unable to find max among args: %v", args)
				}

				return maxObj, nil
			},
		},
	}...)
}
