package grug

import (
	"errors"
	"fmt"
)

// Command represents an invokable Grug command and holds its activators and execution plan
type Command struct {
	Name        string         `yaml:"name"`       // Descriptive name of the command
	Description string         `yaml:"desc"`       // Description of the command
	Activators  []string       `yaml:"activators"` // A list of ways to invoke the command
	Plan        ActionSequence `yaml:"plan"`       // The action sequence to perform when the command is invoked
}

// ConstructActivatorMap takes a list of commands and constructs a map from each command's activators to their respective command
// Returns an error if two commands have conflicting activators
func (g *GrugSession) constructActivatorMap() error {
	g.ActivatorMap = make(map[string]Command)
	for _, c := range g.Commands {
		for _, a := range c.Activators {
			if oc, present := g.ActivatorMap[a]; present {
				return errors.New(fmt.Sprint("commands ", oc.Name, " and ", c.Name, " have a conflicting activator \"", a, "\""))
			}
			g.ActivatorMap[a] = c
		}
	}
	return nil
}
