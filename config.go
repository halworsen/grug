package grug

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type GrugConfig struct {
	Name     string   `yaml:"name"`     // name of the bot
	Token    string   `yaml:"token"`    // token for the bot
	Invoker  string   `yaml:"invoker"`  // the command prefix that invokes grug
	Commands []string `yaml:"commands"` // a list of paths to command config files
	Verbose  bool     `yaml:"verbose"`  // verbose logging
}

// Loads the master grug config into the grug session
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

// Loads the commands specified by the master config
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
