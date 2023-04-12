package operations

import (
	"bufio"
	"fmt"
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


func ListAppsToBeInstalled(appList *[]string)  error {
	if len(*appList) == 0 {
		fmt.Println("No apps will be installed because the list of app is empty. Exiting...")
		os.Exit(0)
	}
	fmt.Println("The following apps will be installed:")
	for _, app := range *appList {
		fmt.Printf("- %s\n", app)
	}
	return nil
}


func AddAppsToList(appList *[]string) ([]string, error) {
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

func RemoveAppsFromList(appList *[]string) ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write the name of the apps you want to remove (separate by space): ")
	text, _ := reader.ReadString('\n')
	apps := strings.TrimSpace(text)
	if apps != "" {
		removedApps := strings.Split(apps, " ")
		_, err := IsElementInSlice(*appList, apps)
		if  err != nil {
			panic(err)
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
	ListAppsToBeInstalled(appList)
	return *appList, nil
}
