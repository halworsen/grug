package grug

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// Creates a file watcher and adds all config files to it
func (g *GrugSession) SetupLiveReloadWatcher() {
	g.Log(logInfo, "Starting live reload watchdog...")
	masterWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		g.Log(logFatal, "Failed to start live reload master config watcher -", err)
	}
	cmdWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		g.Log(logFatal, "Failed to start live reload command config watcher -", err)
	}

	g.Log(logInfo, "Adding master config to master config watcher:", g.Config.ConfigPath)
	err = masterWatcher.Add(g.Config.ConfigPath)
	if err != nil {
		g.Log(logFatal, "Failed to add master config to watcher -", err)
	}

	for _, f := range g.Config.Commands {
		g.Log(logInfo, "Adding command config to command config watcher:", f)
		err = cmdWatcher.Add(f)
		if err != nil {
			g.Log(logFatal, "Failed to add command config to watcher -", err)
		}
	}

	// The monitoring goroutines are responsible for closing their watchers
	go g.MonitorConfigFiles(masterWatcher, g.MasterConfigReloadHandler)
	go g.MonitorConfigFiles(cmdWatcher, g.CommandConfigReloadHandler)
}

// Indefinitely reads watcher events and triggers live reloads as necessary
func (g *GrugSession) MonitorConfigFiles(watcher *fsnotify.Watcher, reloadHandler func(string)(error)) {
	defer watcher.Close()
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op & fsnotify.Write != fsnotify.Write {
				continue
			}

			err := reloadHandler(event.Name)
			if err != nil {
				g.Log(logFatal, "Config reload error:", err)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			g.Log(logError, "Config watcher error:", err)
		}
	}
}

func (g *GrugSession) MasterConfigReloadHandler(fileName string) error {
	g.Log(logInfo, "Master config changes detected. Reloading master config...")
	oldCommands := g.Config.Commands
	err := g.loadMasterConfig(fileName)
	if err != nil {
		return err
	}

	commandsChanged := false
	if len(oldCommands) != len(g.Config.Commands) {
		commandsChanged = true
	}

	if !commandsChanged {
		for i, v := range g.Config.Commands {
			if oldCommands[i] != v {
				commandsChanged = true
				break
			}
		}
	}

	if commandsChanged {
		g.Log(logInfo, "Command list update detected, reloading commands...")
		err = g.loadCommands()
		if err != nil {
			return err
		}
	}

	g.Log(logInfo, "Reconstructing activator map...")
	err = g.constructActivatorMap()
	if err != nil {
		return err
	}

	g.Log(logInfo, "Successfully reloaded master config")
	return nil
}

func (g *GrugSession) CommandConfigReloadHandler(fileName string) error {
	g.Log(logInfo, fmt.Sprint("Command config changes detected in \"", fileName, "\", reloading..."))
	cmd, err := g.loadCommand(fileName)
	if err != nil {
		return err
	}

	// Replace the old command with the new one
	for i, c := range g.Commands {
		if c.ConfigPath == fileName {
			oldTail := g.Commands[i+1:]
			g.Commands = append(g.Commands[:i], cmd)
			g.Commands = append(g.Commands, oldTail...)
		}
	}

	g.Log(logInfo, "Reconstructing activator map...")
	err = g.constructActivatorMap()
	if err != nil {
		return err
	}

	g.Log(logInfo, fmt.Sprint("Successfully reloaded command config from \"", fileName, "\""))
	return nil
}