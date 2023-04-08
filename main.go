package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var defaultApps = []string{"git", "go", "vim"}

func main() {
	var appList []string
	
	// Check if Homebrew is installed and install if not. Throw error if Homebrew installation fails.
	installHomebrew() 

	// Parse flags
	addApps := flag.String("add", "", "Add additional apps (separate by space)")
	removeApps := flag.String("remove", "", "Remove apps (separate by space)")
	installAll := flag.Bool("install-all", false, "Install all apps")
	flag.Parse()

	// If -add flag is set, add apps to default app list
	addAppsHandler(&defaultApps, *addApps)		

	// If -remove flag is set, remove apps from default app list
	removeAppsHandler(&defaultApps, *removeApps)

	// If -install-all flag is set, use default app list
	if *installAll {
		appList = defaultApps
	} else {
		// Otherwise, prompt the user to confirm which apps to install
		listAppsToBeInstalled(defaultApps)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Would you like to install these apps? (y/n): ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(text)) == "y" {
			appList = defaultApps
		} else {
			// Otherwise, allow the user to add or remove apps
			appList = append(appList, defaultApps...)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Do you want to add or remove apps from the list above? (add/remove): ")
			text, _ := reader.ReadString('\n')
			switch strings.TrimSpace(strings.ToLower(text)) {
			case "add":
				appList, err := addAppsToList(&appList)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("\nAdditional apps to be installed: %s \n", *addApps)
				fmt.Printf("\nApp list: %s \n", appList)
			case "remove":
				appList, err := removeAppsFromList(&appList)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("\nRemoved the following app(s): %s \n", *addApps)
				fmt.Printf("\nApp list: %s \n", appList)
			default:
				fmt.Println("Invalid flag. No apps will be installed.")
				os.Exit(0)
			}
		}
	}
	installSelectedApps(appList)
}



func installHomebrew() {
	// Check if Homebrew is already installed
	_, err := exec.LookPath("brew")
	if err == nil {
		fmt.Println("Homebrew is already installed.")
		return
	}

	// Install Homebrew
	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Homebrew: %v", err)
	}
	fmt.Println("Homebrew has been installed.")
}

func addAppsHandler(appList *[]string, addApps string) {
	if addApps != "" {
		fmt.Printf("\nAdditional apps to be installed: %s \n", addApps)
		defaultApps = append(defaultApps, strings.Split(addApps, " ")...)
	}
}

func removeAppsHandler(appList *[]string, removeApps string) {
	if removeApps != "" && isElementInSlice(defaultApps, removeApps) {
		fmt.Printf("\nRemoved the following app(s): %s \n", removeApps)
		removedApps := strings.Split(removeApps, " ")
		for _, app := range removedApps {
			for i, a := range defaultApps {
				if a == app {
					defaultApps = append(defaultApps[:i], defaultApps[i+1:]...)
					break
				}
			}
		}
	}
}

func addAppsToList(appList *[]string) ([]string, error ){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write the name of the apps you want to add (separate by space): ")
	text, _ := reader.ReadString('\n')
	apps := strings.TrimSpace(text)
	if apps != "" {
		addedApps := strings.Split(apps, " ")
		*appList = append(*appList, addedApps...)
	}
	return *appList, nil
}

func isElementInSlice(slice []string, target string) bool {
	var notFound []string
	for _, s := range strings.Split(target, " ") {
		found := false
		for _, e := range slice {
			if e == s {
				found = true
				break
			}
		}
		if !found {
			notFound = append(notFound, s)
		}
	}
	if len(notFound) == 0 {
		return true
	}
	fmt.Printf("\nApp(s) %s is not in the list of apps to be installed. Exiting...\n\n", notFound)
	return false
}


func removeAppsFromList(appList *[]string) ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write the name of the apps you want to remove (separate by space): ")
	text, _ := reader.ReadString('\n')
	apps := strings.TrimSpace(text)
	if apps != "" {
		removedApps := strings.Split(apps, " ")
		if !isElementInSlice(*appList, apps) {
			os.Exit(0)
		}
		for _, app := range removedApps {
			for i, a := range *appList {
				if a == app {
					*appList = append((*appList)[:i], (*appList)[i+1:]...)
					break
				}
			}
		}
	}
	listAppsToBeInstalled(*appList)
	return *appList, nil
}

func installSelectedApps(appList []string)  error {
	start := time.Now()
	var wg sync.WaitGroup
	for _, app := range appList {
		wg.Add(1)
		go func(app string) {
			defer wg.Done()
			fmt.Printf("Installing %s...\n", app)
			if err := runCommand("brew", "install", app); err != nil {
				log.Fatal(err)
			}
		}(app)
	}
	wg.Wait()
	fmt.Printf("All apps have been installed in %s\n", time.Since(start))
	return nil
}

func listAppsToBeInstalled(appList []string) error {
	if len(appList) == 0 {
		fmt.Println("No apps will be installed because the list of app is empty. Exiting...")
		os.Exit(0)
	}
	fmt.Println("The following apps will be installed:")
	for _, app := range appList {
		fmt.Printf("- %s\n", app)
	}
	return nil
}


func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
