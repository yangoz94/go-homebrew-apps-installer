package flags

import (
	"errors"
	"fmt"
	"log"
	"ogi/pkg/operations"
	"os"
	"strings"
)

func AddAppsHandler(appList *[]string, addApps string) error {
    fmt.Printf("\nAdditional apps to be installed: %s \n", addApps)
    newApps := strings.Split(addApps, " ")
    for _, app := range newApps {
        if !contains(*appList, app) {
            *appList = append(*appList, app)
        }
	}
    return nil
}

func contains(appList []string, app string) bool {
    for _, a := range appList {
        if a == app {
            return true
        }
    }
    return false
}

func RemoveAppsHandler(appList *[]string, removeApps string) error {
	if removeApps == "" {
		return errors.New("removeApps cannot be empty")
	}
	found, err := operations.IsElementInSlice(*appList, removeApps)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("%s not found in appList", removeApps)
	}
	log.Printf("\nRemoved the following app(s): %s \n", removeApps)
	removedApps := strings.Split(removeApps, " ")
	for _, app := range removedApps {
		for i, a := range *appList {
			if a == app {
				*appList = append((*appList)[:i], (*appList)[i+1:]...)
				break
			}
		}
	}
	return nil
}

type UserInputReader interface {
	ReadString(delim byte) (string, error)
}

type Internals interface {
	ListAppsToBeInstalled(appList *[]string)
	AddAppsToList(appList *[]string) ([]string, error)
	RemoveAppsFromList(appList *[]string) ([]string, error)
}
func InstallAllHandler(appList *[]string, installAll *bool, addApps *string, removeApps *string, reader UserInputReader, internals Internals) error {
	internals.ListAppsToBeInstalled(appList)
	log.Println("Would you like to install these apps? (y/n) Type (n) to add/remove apps from the given list. ")
	text, _ := reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(text)) == "y" {
		*installAll = true
	} else {
		log.Print("Do you want to add or remove apps from the list above? (add/remove): ")
		text, _ := reader.ReadString('\n')
		switch strings.TrimSpace(strings.ToLower(text)) {
		case "add":
			appList, err := internals.AddAppsToList(appList)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("\nAdditional apps to be installed: %s \n", *addApps)
			log.Printf("\nUpdated to be installed: %s \n", appList)
		case "remove":
			appList, err := internals.RemoveAppsFromList(appList)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("\nRemoved the following app(s): %s \n", *addApps)
			log.Printf("\nUpdated to be installed: %s \n", appList)
		default:
			log.Println("Invalid flag. No apps will be installed.")
			os.Exit(0)
		}
	}
	return nil
}