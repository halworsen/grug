package grug

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GrugConfig holds the master configuration values for a Grug session
type GrugConfig struct {
	ConfigPath string                       // Path to the file this config was loaded from
	Name       string   `yaml:"name"`       // Name of the bot
	Token      string   `yaml:"token"`      // Discord Bot Token to use
	Invoker    string   `yaml:"invoker"`    // The command prefix that invokes grug
	Commands   []string `yaml:"commands"`   // A list of paths to command config files
	Verbose    bool     `yaml:"verbose"`    // Whether or not to use verbose logging
	HardError  bool     `yaml:"hardError"`  // If set to true, Grug will crash on panics in commands
	LiveReload bool     `yaml:"liveReload"` // Whether or not to use live reloading of command configurations
}

// LoadMasterConfig loads the master grug config into the grug session
func (g *GrugSession) loadMasterConfig(cfgPath string) error {
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &g.Config)
	if err != nil {
		return err
	}
	g.Config.ConfigPath = cfgPath

	return nil
}

// LoadCommands loads the commands specified by the master config
func (g *GrugSession) loadCommands() error {
	g.Commands = nil
	g.Commands = make([]Command, 0)
	for _, f := range g.Config.Commands {
		g.Log(logInfo, fmt.Sprint("Loading command from \"", f, "\""))
		cmd, err := g.loadCommand(f)
		if err != nil {
			return err
		}

		g.Commands = append(g.Commands, cmd)
		g.Log(logInfo, fmt.Sprint("Loaded command \"", cmd.Name, "\" (plan length: ", len(cmd.Plan), ", activators: ", len(cmd.Activators), ")"))
	}

	return nil
}

func (g *GrugSession) loadCommand(filePath string) (Command, error) {
	cmd := Command{ConfigPath: filePath}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cmd, err
	}
	
	err = yaml.Unmarshal(data, &cmd)
	if err != nil {
		return cmd, err
	}

	return cmd, nil
}