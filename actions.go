package grug

import (
	"errors"
	"fmt"
)

// Named action that performs some task
type Action struct {
	Name string                                                  // Name of the action
	Exec func(*GrugSession, ...interface{}) (interface{}, error) // Function that executes the action
}

// YAML-(un)marshallable struct for storing the name of and arguments for an action
type ActionActivator struct {
	ActionName  string             `yaml:"action"` // Name of the action to execute
	Arguments   []interface{}      `yaml:"args"`   // Arguments taken by the action and their type
	Store       string             `yaml:"store"`  // If set, the result of the action is stored in a field with this name
	Conditional *ConditionalAction `yaml:"if"`     // If set, this activator is considered an conditional and will instead alter flow of the sequence
}

// Action sequences are a list of action activations to be performed in sequential order
type ActionSequence []ActionActivator

// A conditional. Performs no action, but determines which action sequence is taken after evaluation
type ConditionalAction struct {
	ActionName    string         `yaml:"condition"` // Name of the action performing the conditional
	Arguments     []interface{}  `yaml:"args"`      // Arguments taken by the conditional action
	TrueSequence  ActionSequence `yaml:"true"`      // The action sequence to perform when the condition evaluates to true
	FalseSequence ActionSequence `yaml:"false"`     // The action sequence to perform when the condition evaluates to false
}

var AllActions []Action

// Takes a list of actions and constructs a map from each action's name to their respective action struct
func (g *GrugSession) ConstructActionMap() {
	g.ActionMap = make(map[string]Action)
	for _, a := range g.Actions {
		g.ActionMap[a.Name] = a
	}
}

func (g *GrugSession) PerformStep(activator ActionActivator, userArgs []string) error {
	// check which action sequence should be performed
	if activator.Conditional != nil {
		// invalid configurations
		action, present := g.ActionMap[activator.Conditional.ActionName]
		if !present {
			return errors.New(fmt.Sprint("bad conditional action name: ", activator.ActionName, ""))
		}

		args, err := ParseArgs(activator.Conditional.Arguments, userArgs)
		if err != nil {
			return errors.New(fmt.Sprint("failed to parse arguments for action ", activator.ActionName, " - ", err))
		}

		result, err := action.Exec(g, args...)
		if err != nil {
			return errors.New(fmt.Sprint("failed to execute conditional action (name: ", activator.Conditional.ActionName, ") - ", err))
		}
		// check that we actually got a bool
		resultBool, ok := result.(bool)
		if !ok {
			return errors.New("conditional action returned non-boolean result")
		}

		// perform the true/false sequence depending on the result
		var newSeq ActionSequence
		newSeq = activator.Conditional.FalseSequence
		if resultBool {
			newSeq = activator.Conditional.TrueSequence
		}

		// it's okay to have empty action sequences in conditionals
		for _, newActivator := range newSeq {
			err := g.PerformStep(newActivator, userArgs)
			if err != nil {
				return err
			}
		}
	} else {
		// invalid configurations
		action, present := g.ActionMap[activator.ActionName]
		if !present {
			return errors.New(fmt.Sprint("bad action name: ", activator.ActionName, ""))
		}

		args, err := ParseArgs(activator.Arguments, userArgs)
		if err != nil {
			return errors.New(fmt.Sprint("failed to parse arguments for action ", activator.ActionName, " - ", err))
		}

		result, err := action.Exec(g, args...)
		if err != nil {
			return errors.New(fmt.Sprint("failed to execute action (name: ", activator.ActionName, ") - ", err))
		}
		// Store the result of this step
		if activator.Store != "" {
			err = StoreArg(activator.Store, result)
			if err != nil {
				return errors.New(fmt.Sprint("failed to store result of action ", activator.ActionName, " - ", err))
			}
		}
	}

	return nil
}
