package grug

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GrugConfig holds the master configuration values for a Grug session
type GrugConfig struct {
	Name      string   `yaml:"name"`      // Name of the bot
	Token     string   `yaml:"token"`     // Discord Bot Token to use
	Invoker   string   `yaml:"invoker"`   // The command prefix that invokes grug
	Commands  []string `yaml:"commands"`  // A list of paths to command config files
	Verbose   bool     `yaml:"verbose"`   // Whether or not to use verbose logging
	HardError bool     `yaml:"hardError"` // If set to true, Grug will crash on panics in commands
}

// LoadMasterConfig loads the master grug config into the grug session
func (g *GrugSession) LoadMasterConfig(cfgPath string) error {
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &g.Config)
	if err != nil {
		return err
	}

	return nil
}

// LoadCommands loads the commands specified by the master config
func (g *GrugSession) LoadCommands() error {
	for _, f := range g.Config.Commands {
		g.Log(logInfo, fmt.Sprint("Loading command from \"", f, "\""))
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		cmd := Command{}
		err = yaml.Unmarshal(data, &cmd)
		if err != nil {
			return err
		}

		g.Commands = append(g.Commands, cmd)
		g.Log(logInfo, fmt.Sprint("Loaded command \"", cmd.Name, "\" (plan length: ", len(cmd.Plan), ", activators: ", len(cmd.Activators), ")"))
	}

	return nil
}
