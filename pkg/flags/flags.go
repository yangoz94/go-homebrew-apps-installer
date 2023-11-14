package flags

import (
	"errors"
	"fmt"
	"log"
	"ogi/pkg/internals"
	"ogi/pkg/operations"
	"strings"
)

func AddAppsHandler(appList *[]string, addApps string) error {
	fmt.Printf("Additional apps to be installed: %s ", addApps)
	newApps := strings.Split(addApps, " ")
	for _, app := range newApps {
		if !operations.Contains(*appList, app) {
			*appList = append(*appList, app)
		}
	}
	return nil
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

	for {
		log.Println("Would you like to install these apps? (y/n) Type (n) to add/remove apps from the given list.")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(strings.ToLower(text))

		switch text {
		case "y":
			*installAll = true
			return nil
		case "n":
			for {
				log.Print("Do you want to add or remove apps from the list above? (add/remove): ")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(strings.ToLower(text))

				switch text {
				case "add", "remove":
					apps := operations.ReadAppList()
					if len(apps) == 0 {
						log.Println("No apps specified. Please enter the names of the app(s)")
						continue
					}
					var err error
					if text == "add" {
						*appList, err = internals.AddAppsToList(appList, apps)
						log.Printf("Additional apps to be installed: %s \n", apps)
					} else {
						for _, app := range strings.Split(apps, " ") {
							if !operations.Contains(*appList, app) {
								log.Fatal("App not found in the list of apps to be installed")
							}
						}
						*appList, err = internals.RemoveAppsFromList(appList, apps)
					}

					if err != nil {
						log.Fatal(err)
					}
					log.Printf("Updated to be installed: %s \n", *appList)
					return nil
				default:
					log.Println("Invalid command. Please enter 'add' or 'remove'.")
				}
			}
		default:
			log.Println("Invalid command. Please enter 'y' or 'n'.")
		}
	}

}
