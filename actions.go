package grug

import (
	"errors"
	"fmt"
)

// Action is any named task
type Action struct {
	Name string                                                  // Name of the action
	Exec func(*GrugSession, ...interface{}) (interface{}, error) // Function that executes the action
}

// ActionActivator is a YAML-(un)marshallable struct for storing the name of and arguments for an action
type ActionActivator struct {
	ActionName  string             `yaml:"action"`        // Name of the action to execute
	Arguments   []interface{}      `yaml:"args"`          // Arguments taken by the action and their type
	Store       string             `yaml:"store"`         // If set, the result of the action is stored in a field with this name
	FailurePlan *ActionSequence    `yaml:"failurePlan"`   // An action sequence to perform if this step fails during execution
	HaltOnFail  bool               `yaml:"haltOnFailure"` // Whether or not to halt command execution on failure of this step
	Conditional *ConditionalAction `yaml:"if"`            // If set, this activator is considered an conditional and will instead alter flow of the sequence
}

// ActionSequence is a list of action activations to be performed in sequential order
type ActionSequence []ActionActivator

// ConditionalAction represents a conditional part of an action sequence. Performs no action, but determines which action sequence is taken after evaluation
type ConditionalAction struct {
	ActionName    string         `yaml:"condition"` // Name of the action performing the conditional
	Arguments     []interface{}  `yaml:"args"`      // Arguments taken by the conditional action
	TrueSequence  ActionSequence `yaml:"true"`      // The action sequence to perform when the condition evaluates to true
	FalseSequence ActionSequence `yaml:"false"`     // The action sequence to perform when the condition evaluates to false
}

// AllActions holds all actions that can be used by Grug
var AllActions []Action

// ConstructActionMap takes a list of actions and constructs a map from each action's name to their respective action struct
func (g *GrugSession) ConstructActionMap() {
	g.ActionMap = make(map[string]Action)
	for _, a := range g.Actions {
		g.ActionMap[a.Name] = a
	}
}

// PerformAction performs an action given an activator for the function and any user supplied arguments
func (g *GrugSession) PerformAction(activator ActionActivator, userArgs []string) error {
	actionName, args := activator.ActionName, activator.Arguments
	if activator.Conditional != nil {
		actionName = activator.Conditional.ActionName
		args = activator.Conditional.Arguments
	}

	// invalid configurations
	action, ok := g.ActionMap[actionName]
	if !ok {
		return errors.New(fmt.Sprint("bad action name: ", activator.ActionName, ""))
	}

	args, err := ParseArgs(args, userArgs)
	if err != nil {
		return errors.New(fmt.Sprint("failed to parse arguments for action ", activator.ActionName, " - ", err))
	}

	result, err := action.Exec(g, args...)
	if err != nil {
		// The step failed, so perform the failure action sequence
		if activator.FailurePlan != nil {
			for fStep, newActivator := range *activator.FailurePlan {
				err := g.PerformAction(newActivator, userArgs)
				if err != nil {
					g.Log(logError, fmt.Sprint("Failed to execute failure step ", fStep, " - ", err))
				}
			}
		}
		return errors.New(fmt.Sprint("failed to execute action ", activator.ActionName, " - ", err))
	}

	if activator.Conditional == nil {
		// Store the result of this step
		if activator.Store != "" {
			err = StoreArg(activator.Store, result)
			if err != nil {
				return errors.New(fmt.Sprint("failed to store result of action ", activator.ActionName, " - ", err))
			}
		}
	} else {
		// Check that we actually got a bool
		resultBool, ok := result.(bool)
		if !ok {
			return errors.New("conditional action returned non-boolean result")
		}

		// Perform the true/false sequence depending on the result
		var newSeq ActionSequence
		newSeq = activator.Conditional.FalseSequence
		if resultBool {
			newSeq = activator.Conditional.TrueSequence
		}

		// It's okay to have empty action sequences in conditionals
		for _, newActivator := range newSeq {
			err := g.PerformAction(newActivator, userArgs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
