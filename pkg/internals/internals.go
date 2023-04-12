package internals

import "ogi/pkg/operations"

type DefaultInternals struct{}

func (d *DefaultInternals) ListAppsToBeInstalled(appList *[]string) {
	operations.ListAppsToBeInstalled(appList)
}

func (d *DefaultInternals) AddAppsToList(appList *[]string) ([]string, error) {
	return operations.AddAppsToList(appList)
}

func (d *DefaultInternals) RemoveAppsFromList(appList *[]string) ([]string, error) {
	return operations.RemoveAppsFromList(appList)
}


