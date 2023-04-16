package operations

import (
	"reflect"
	"strings"
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
    // Define the test cases as a slice of structs:
    testCases := []struct {
        name          string
        appList       []string
        appsToAdd     string
        expectedList  []string
        expectedError error
    }{
        {
            name:          "Add to empty list",
            appList:       []string{},
            appsToAdd:     "app1 app2 app3",
            expectedList:  []string{"app1", "app2", "app3"},
            expectedError: nil,
        },
		{
            name:          "Add a single app to non-empty list",
            appList:       []string{"app1", "app2"},
            appsToAdd:     "app3",
            expectedList:  []string{"app1", "app2", "app3"},
            expectedError: nil,
        },
        {
            name:          "Add multiple apps to a non-empty list",
            appList:       []string{"app1", "app2"},
            appsToAdd:     "app3 app4",
            expectedList:  []string{"app1", "app2", "app3", "app4"},
            expectedError: nil,
        },
        {
            name:          "Add empty string",
            appList:       []string{"app1", "app2"},
            appsToAdd:     "",
            expectedList:  []string{"app1", "app2"},
            expectedError: nil,
        },
    }

    // Loop through the test cases and run each one:
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := AddAppsToList(&tc.appList, tc.appsToAdd)

            // Check that the app list was modified correctly:
            if !reflect.DeepEqual(result, tc.expectedList) {
                t.Errorf("Expected %v, but got %v", tc.expectedList, result)
            }

            // Check that the error is nil:
            if err != tc.expectedError {
                t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
            }
        })
    }
}

func TestRemoveAppsFromList(t *testing.T) {
    // Define the test cases as a slice of structs:
    testCases := []struct {
        name          string
        appList       []string
        appsToRemove  string
        expectedList  []string
        expectedError string
    }{
        {
            name:          "Remove one app",
            appList:       []string{"app1", "app2", "app3"},
            appsToRemove:  "app1",
            expectedList:  []string{"app2", "app3"},
            expectedError: "",
        },
        {
            name:          "Remove multiple apps",
            appList:       []string{"app1", "app2", "app3"},
            appsToRemove:  "app1 app2",
            expectedList:  []string{"app3"},
            expectedError: "",
        },
        {
            name:          "Remove all apps",
            appList:       []string{"app1", "app2", "app3"},
            appsToRemove:  "app1 app2 app3",
            expectedList:  []string{},
            expectedError: "",
        },
        {
            name:          "Remove a non-existent app",
            appList:       []string{"app1", "app2", "app3"},
            appsToRemove:  "app4",
            expectedList:  []string{"app1", "app2", "app3"},
            expectedError: "App(s) app4 is/are not in the list of apps to be installed",
        },
        {
            name:          "Remove empty string",
            appList:       []string{"app1", "app2", "app3"},
            appsToRemove:  "",
            expectedList:  []string{"app1", "app2", "app3"},
            expectedError: "",
        },
    }

    // Loop through the test cases and run each one:
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create a copy of the app list for each test case:
            list := make([]string, len(tc.appList))
            copy(list, tc.appList)

            // Call the RemoveAppsFromList function and check the result:
            result, err := RemoveAppsFromList(&list, tc.appsToRemove)

            if !reflect.DeepEqual(result, tc.expectedList) {
                t.Errorf("Expected %v, but got %v", tc.expectedList, result)
            }

            if tc.expectedError != "" && err == nil {
                t.Errorf("Expected error containing '%s', but got no error", tc.expectedError)
            }

            if tc.expectedError == "" && err != nil {
                t.Errorf("Expected no error, but got '%v'", err)
            }

            if tc.expectedError != "" && err != nil && !strings.Contains(err.Error(), tc.expectedError) {
                t.Errorf("Expected error containing '%s', but got '%v'", tc.expectedError, err)
            }
        })
    }
}