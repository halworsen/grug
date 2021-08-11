package grug

import (
	"context"
	"strconv"

	"github.com/halworsen/grug/util"
)

func init() {
	AllActions = append(AllActions, []Action{
		{
			Name: "==",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				return (args[0] == args[1]), nil
			},
		},
		{
			Name: "not",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				b, err := strconv.ParseBool(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				return !b, nil
			},
		},
		{
			Name: "int>",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a > b), nil
			},
		},
		{
			Name: "int>=",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a >= b), nil
			},
		},
		{
			Name: "int<",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a > b), nil
			},
		},
		{
			Name: "int<=",
			Exec: func(g *GrugSession, ctx context.Context, args ...interface{}) (interface{}, error) {
				a, err := strconv.Atoi(util.Atostr(args[0]))
				if err != nil {
					return nil, err
				}
				b, err := strconv.Atoi(util.Atostr(args[1]))
				if err != nil {
					return nil, err
				}
				return (a >= b), nil
			},
		},
	}...)
}
