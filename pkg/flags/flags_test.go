package flags

import (
	"reflect"
	"testing"
)

func TestAddAppsHandler(t *testing.T) {
	tests := []struct {
		name     string
		appList  []string
		addApps  string
		expected []string
	}{
		{
			name:     "Test adding a single app to an empty app list",
			appList:  []string{},
			addApps:  "app1",
			expected: []string{"app1"},
		},
		{
			name:     "Test adding multiple apps to an empty app list",
			appList:  []string{},
			addApps:  "app1 app2",
			expected: []string{"app1", "app2"},
		},
		{
			name:     "Test adding a single app to a non-empty app list",
			appList:  []string{"app1"},
			addApps:  "app2",
			expected: []string{"app1", "app2"},
		},
		{
			name:     "Test adding multiple apps to a non-empty app list",
			appList:  []string{"app1"},
			addApps:  "app2 app3",
			expected: []string{"app1", "app2", "app3"},
		},
		{
			name:     "Test adding an app that already exists in the app list",
			appList:  []string{"app1"},
			addApps:  "app1",
			expected: []string{"app1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := AddAppsHandler(&test.appList, test.addApps)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.appList, test.expected) {
				t.Errorf("Expected app list %v, got %v", test.expected, test.appList)
			}
		})
	}
}

func TestRemoveAppsHandler(t *testing.T) {
	tests := []struct {
		name       string
		appList    []string
		removeApps string
		expected   []string
	}{
		{
			name:       "Test removing a single app from an app list with only one app",
			appList:    []string{"app1"},
			removeApps: "app1",
			expected:   []string{},
		},
		{
			name:       "Test removing multiple apps from an app list with multiple apps",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "app1 app3",
			expected:   []string{"app2"},
		},
		{
			name:       "Test removing a single app from an app list with multiple apps",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "app2",
			expected:   []string{"app1", "app3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RemoveAppsHandler(&test.appList, test.removeApps)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.appList, test.expected) {
				t.Errorf("Expected app list %v, got %v", test.expected, test.appList)
			}
		})
	}
}

