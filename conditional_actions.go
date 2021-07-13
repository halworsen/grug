package grug

import "strconv"

func init() {
	AllActions = append(AllActions, []Action{
		{
			Name: "==",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				return (args[0] == args[1]), nil
			},
		},
		{
			Name: "not",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				b, err := strconv.ParseBool(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				return !b, nil
			},
		},
		{
			Name: "int>",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a > b), nil
			},
		},
		{
			Name: "int>=",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a >= b), nil
			},
		},
		{
			Name: "int<",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a > b), nil
			},
		},
		{
			Name: "int<=",
			Exec: func(g *GrugSession, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a >= b), nil
			},
		},
	}...)
}
