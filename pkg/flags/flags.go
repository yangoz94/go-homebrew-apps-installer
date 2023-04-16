package flags

import (
	"errors"
	"fmt"
	"log"
	"ogi/pkg/internals"
	"ogi/pkg/operations"
	"os"
	"strings"
)

func AddAppsHandler(appList *[]string, addApps string) error {
    fmt.Printf("Additional apps to be installed: %s ", addApps)
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
	log.Printf("\nRemoved the following app(s): %s", removeApps)
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


func InstallAllHandler(appList *[]string, installAll *bool, addApps *string, removeApps *string, reader UserInputReader, internals internals.Internals) error {
	internals.ListAppsToBeInstalled(appList)

	log.Println("Would you like to install these apps? (y/n) Type (n) to add/remove apps from the given list.")
	text, _ := reader.ReadString('\n')

	if strings.TrimSpace(strings.ToLower(text)) == "y" {
		*installAll = true
	} else {
		log.Print("Do you want to add or remove apps from the list above? (add/remove): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(strings.ToLower(text))

		switch text {
		case "add", "remove":
			apps := operations.ReadAppList()
			var err error
			if text == "add" {
				*appList, err = internals.AddAppsToList(appList, apps)
				log.Printf("Additional apps to be installed: %s \n", apps)
			} else {
				*appList, err = internals.RemoveAppsFromList(appList, apps)
				log.Printf("Removed the following app(s): %s \n", apps)
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Updated to be installed: %s \n", *appList)
		default:
			log.Println("Invalid flag. No apps will be installed.")
			os.Exit(0)
		}
	}

	return nil
}
