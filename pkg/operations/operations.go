package operations

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func IsElementInSlice(slice []string, target string) (bool, error) {
	targets := strings.Split(target, " ")
	for _, t := range targets {
		for _, element := range slice {
			if element == t {
				return true, nil
			}
		}
	}
	return false, fmt.Errorf("App(s) %s is/are not in the list of apps to be installed", target)
}

func ListAppsToBeInstalled(appList *[]string) error {
	if len(*appList) == 0 {
		return errors.New("no apps will be installed because the list of app is empty - exiting")
	}
	log.Println("The following apps will be installed:")
	for _, app := range *appList {
		fmt.Printf("- %s\n", app)
	}
	return nil
}

func Contains(appList []string, app string) bool {
	for _, a := range appList {
		if a == app {
			return true
		}
	}
	return false
}

func AddAppsToList(appList *[]string, appsToAdd string) ([]string, error) {
	if appsToAdd != "" {
		addedApps := strings.Split(appsToAdd, " ")
		for _, app := range addedApps {
			if !Contains(*appList, app) {
				*appList = append(*appList, app)
			}
		}
	}
	ListAppsToBeInstalled(appList)
	return *appList, nil
}

func RemoveAppsFromList(appList *[]string, appsToRemove string) ([]string, error) {
	if appsToRemove == "" {
		return *appList, nil
	}
	removedApps := strings.Split(appsToRemove, " ")
	notFound := []string{}
	for _, app := range removedApps {
		found := false
		for i, a := range *appList {
			if a == app {
				*appList = append((*appList)[:i], (*appList)[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			notFound = append(notFound, app)
		}
	}
	if len(notFound) > 0 {
		return *appList, fmt.Errorf("App(s) %s is/are not in the list of apps to be installed", strings.Join(notFound, ", "))
	}
	ListAppsToBeInstalled(appList)
	return *appList, nil
}

func ReadAppList() string {
	fmt.Println("Write the name of the apps (separate by space): ")
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	apps := strings.TrimSpace(text)
	return apps
}
