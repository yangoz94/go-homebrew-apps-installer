package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"ogi/config"
	"ogi/pkg/flags"
	"ogi/pkg/installers"
	"ogi/pkg/internals"
)

func main() {
	//remove timestamp from logs for better readability and console interaction(tests will still work if this is removed)
	log.SetFlags(0)

	// Check if homebrew is installed.Skip if installed, install if not. Throws error if installation fails.
	installers.InstallHomebrew()

	// flags to customize the list of apps to install
	addApps := flag.String("add", "", "Add additional apps (separate by space)")
	removeApps := flag.String("remove", "", "Remove apps (separate by space)")
	installAll := flag.Bool("installAll", false, "Install all apps")
	flag.Parse()

	//handle flags
	if *addApps != "" {
		err := flags.AddAppsHandler(&config.DEFAULT_APPS, *addApps)
		if err != nil {
			log.Fatalf("Error adding apps: %v", err)
		}
	}

	if *removeApps != "" {
		err := flags.RemoveAppsHandler(&config.DEFAULT_APPS, *removeApps)
		if err != nil {
			log.Fatalf("Error removing apps: %v", err)
		}
	}

	if !*installAll {
		err := flags.InstallAllHandler(&config.DEFAULT_APPS, installAll, addApps, removeApps, bufio.NewReader(os.Stdin), &internals.DefaultInternals{})
		if err != nil {
			log.Fatalf("Error installing apps: %v", err)
		}
	}

	// install apps
	err := installers.InstallSelectedApps(&config.DEFAULT_APPS)
	if err != nil {
		log.Fatalf("Error installing apps: %v", err)
	}
}
