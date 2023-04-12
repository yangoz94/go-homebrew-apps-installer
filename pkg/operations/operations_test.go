// TO-DO add tests for AddAppsToList and RemoveAppsFromList
package operations

import (
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

