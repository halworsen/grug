package grug

// Named action that performs some task
type Action struct {
	Name string                                             // Name of the action
	Exec func(*GrugSession, ...interface{}) (string, error) // Function that executes the action
}

// YAML-(un)marshallable struct for storing the name of and arguments for an action
type ActionActivator struct {
	ActionName string        `yaml:"action"`     // Name of the action to execute
	Arguments  []interface{} `yaml:"args"`       // Arguments taken by the action and their type
	PushResult bool          `yaml:"pushResult"` // Whether or not the result of this step should be pushed onto the stack
}

var AllActions []Action

// Takes a list of actions and constructs a map from each action's name to their respective action struct
func (g *GrugSession) ConstructActionMap() {
	g.ActionMap = make(map[string]Action)
	for _, a := range g.Actions {
		g.ActionMap[a.Name] = a
	}
}
