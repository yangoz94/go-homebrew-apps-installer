package operations

import (
	"reflect"
	"testing"
)

func TestRunCommand(t *testing.T) {
	err := RunCommand("echo", "hello")
	if err != nil {
		t.Errorf("Error running command: %v", err)
	}
}

func TestIsElementInSlice(t *testing.T) {
	slice := []string{"apple", "banana", "orange"}
	found, err := IsElementInSlice(slice, "banana")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !found {
		t.Errorf("Expected element to be in slice")
	}

	found, err = IsElementInSlice(slice, "pear")
	if err == nil {
		t.Errorf("Expected error but got none")
	}
	if found {
		t.Errorf("Expected element to not be in slice")
	}
}

func TestListAppsToBeInstalled(t *testing.T) {
	appList := []string{"wget", "curl"}
	err := ListAppsToBeInstalled(&appList)
	if err != nil {
		t.Errorf("Error listing apps: %v", err)
	}
}

func TestAddAppsToList(t *testing.T) {
	// mock the ReadAppList function to return a fixed string
	ReadAppList = func(appList *[]string) (string, error) {
		return "app1 app2 app3", nil
	}
	// create a test app list
	appList := []string{"app4", "app5"}
	// call the function to add apps to the list
	expected := []string{"app4", "app5", "app1", "app2", "app3"}
	actual, err := AddAppsToList(&appList)
	// check for errors
	if err != nil {
		t.Errorf("AddAppsToList returned an error: %v", err)
	}
	// check if the actual list matches the expected list
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("AddAppsToList returned %v, expected %v", actual, expected)
	}
}

func TestRemoveAppsFromList(t *testing.T) {
	// mock the ReadAppList function to return a fixed string
	ReadAppList = func(appList *[]string) (string, error) {
		return "app2 app4", nil
	}
	// create a test app list
	appList := []string{"app1", "app2", "app3", "app4", "app5"}
	// call the function to remove apps from the list
	expected := []string{"app1", "app3", "app5"}
	actual, err := RemoveAppsFromList(&appList)
	// check for errors
	if err != nil {
		t.Errorf("RemoveAppsFromList returned an error: %v", err)
	}
	// check if the actual list matches the expected list
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RemoveAppsFromList returned %v, expected %v", actual, expected)
	}
}

func TestReadAppList(t *testing.T) {
	// mock the ReadAppList function to return a fixed string
	ReadAppList = func(appList *[]string) (string, error) {
		return "Test", nil
	}
	// create a test app list
	appList := []string{}
	// call the function to read apps from user input
	expected := "Test"
	actual, err := ReadAppList(&appList)
	// check for errors
	if err != nil {
		t.Errorf("ReadAppList returned an error: %v", err)
	}
	// check if the actual string matches the expected string
	if actual != expected {
		t.Errorf("ReadAppList returned %v, expected %v", actual, expected)
	}
}