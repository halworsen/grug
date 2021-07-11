package grug

import (
	"errors"
	"fmt"
)

type Command struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"desc"`
	Activators  []string          `yaml:"activators"`
	Steps       []ActionActivator `yaml:"steps"`
}

// Takes a list of commands and constructs a map from each command's activators to their respective command
// Returns an error if two commands have conflicting activators
func (g *GrugSession) ConstructActivatorMap() error {
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
