package internals

import (
	"ogi/pkg/operations"
)

type DefaultInternals struct{}

func (d *DefaultInternals) ListAppsToBeInstalled(appList *[]string) {
	operations.ListAppsToBeInstalled(appList)
}

func (d *DefaultInternals) AddAppsToList(appList *[]string, apps string) ([]string, error) {
	return operations.AddAppsToList(appList, apps)
}

func (d *DefaultInternals) RemoveAppsFromList(appList *[]string, apps string) ([]string, error) {
	return operations.RemoveAppsFromList(appList, apps)
}

// for flags package
type Internals interface {
	ListAppsToBeInstalled(appList *[]string)
	AddAppsToList(appList *[]string, apps string) ([]string, error)
	RemoveAppsFromList(appList *[]string, apps string) ([]string, error)
}
