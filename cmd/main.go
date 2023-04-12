package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"ogi/pkg/flags"
	"ogi/pkg/installers"
	"ogi/pkg/internals"

)

type Config struct {
	DefaultApps []string `json:"defaultApps"`
}

func main() {
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	defaultApps := config.DefaultApps

	installers.InstallHomebrew()

	addApps := flag.String("add", "", "Add additional apps (separate by space)")
	removeApps := flag.String("remove", "", "Remove apps (separate by space)")
	installAll := flag.Bool("installAll", false, "Install all apps")
	flag.Parse()

	if *addApps != "" {
		err := flags.AddAppsHandler(&defaultApps, *addApps)
		if err != nil {
			log.Fatalf("Error adding apps: %v", err)
		}
	}

	if *removeApps != "" {
		err := flags.RemoveAppsHandler(&defaultApps, *removeApps)
		if err != nil {
			log.Fatalf("Error removing apps: %v", err)
		}
	}

	if !*installAll {
		err := flags.InstallAllHandler(&defaultApps, installAll, addApps, removeApps,bufio.NewReader(os.Stdin), &internals.DefaultInternals{})
		if err != nil {
			log.Fatalf("Error installing apps: %v", err)
		}
	}

	err = installers.InstallSelectedApps(&defaultApps)
	if err != nil {
		log.Fatalf("Error installing apps: %v", err)
	}
}