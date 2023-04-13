package flags

import (
	"bytes"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
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
		wantErr    bool
	}{
		{
			name:       "Test removing a single app from an app list with only one app",
			appList:    []string{"app1"},
			removeApps: "app1",
			expected:   []string{},
			wantErr:    false,
		},
		{
			name:       "Test removing multiple apps from an app list with multiple apps",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "app1 app3",
			expected:   []string{"app2"},
			wantErr:    false,
		},
		{
			name:       "Test removing a single app from an app list with multiple apps",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "app2",
			expected:   []string{"app1", "app3"},
			wantErr:    false,
		},
		{
			name:       "Test removing an empty string from an app list",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "",
			expected:   []string{"app1", "app2", "app3"},
			wantErr:    true,
		},
		{
			name:       "Test removing an app that is not in the app list",
			appList:    []string{"app1", "app2", "app3"},
			removeApps: "app4",
			expected:   []string{"app1", "app2", "app3"},
			wantErr:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RemoveAppsHandler(&test.appList, test.removeApps)
			if (err != nil) != test.wantErr {
				t.Errorf("RemoveAppsHandler() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(test.appList, test.expected) {
				t.Errorf("Expected app list %v, got %v", test.expected, test.appList)
			}
		})
	}
}

type mockReader struct {
	inputs []string
	index  int
}

func (m *mockReader) ReadString(delim byte) (string, error) {
	if m.index >= len(m.inputs) {
		return "", io.EOF
	}
	result := m.inputs[m.index]
	m.index++
	return result, nil
}

type mockInternals struct {
	listAppsToBeInstalled func(appList *[]string)
	addAppsToList         func(appList *[]string) ([]string, error)
	removeAppsFromList    func(appList *[]string) ([]string, error)
}

func (m *mockInternals) ListAppsToBeInstalled(appList *[]string) {
	if m.listAppsToBeInstalled != nil {
		m.listAppsToBeInstalled(appList)
	}
}

func (m *mockInternals) AddAppsToList(appList *[]string) ([]string, error) {
	if m.addAppsToList != nil {
		return m.addAppsToList(appList)
	}
	return nil, nil
}

func (m *mockInternals) RemoveAppsFromList(appList *[]string) ([]string, error) {
	if m.removeAppsFromList != nil {
		return m.removeAppsFromList(appList)
	}
	return nil, nil
}

func TestInstallAllHandler(t *testing.T) {
	tests := []struct {
		name           string
		appList        []string
		readerInputs   []string
		listApps       func(appList *[]string)
		addApps        func(appList *[]string) ([]string, error)
		removeApps     func(appList *[]string) ([]string, error)
		expectedOutput string
	}{
		{
			name:         "Test installing all apps",
			appList:      []string{"app1", "app2", "app3"},
			readerInputs: []string{"y\n"},
			listApps: func(appList *[]string) {
				*appList = []string{"app1", "app2", "app3"}
			},
			expectedOutput: "Would you like to install these apps? (y/n): \n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer func() {
				log.SetOutput(os.Stderr)
			}()
			reader := &mockReader{inputs: test.readerInputs}
			internals := &mockInternals{
				listAppsToBeInstalled: test.listApps,
				addAppsToList:         test.addApps,
				removeAppsFromList:    test.removeApps,
			}
			installAll := false
			addApps := ""
			removeApps := ""
			err := InstallAllHandler(&test.appList, &installAll, &addApps, &removeApps, reader, internals)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			output := buf.String()
			// This is to ignore the timestamp in the output due to the log statement
			re := regexp.MustCompile(`\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} `)
			output = re.ReplaceAllString(output, "")
			if output != test.expectedOutput {
				t.Errorf("Expected output %q, got %q", test.expectedOutput, output)
			}
		})
	}
}
