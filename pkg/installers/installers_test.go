package installers
import (
	"os/exec"
	"testing"
)

func TestInstallHomebrew(t *testing.T) {
	// Test that Homebrew is installed after running the function
	InstallHomebrew()
	_, err := exec.LookPath("brew")
	if err != nil {
		t.Errorf("Homebrew was not installed: %v", err)
	}

	// Test that the function correctly detects when Homebrew is already installed
	InstallHomebrew()
}

func TestInstallSelectedApps(t *testing.T) {
	appList := []string{"wget", "curl"}
	err := InstallSelectedApps(&appList)
	if err != nil {
		t.Errorf("Error installing apps: %v", err)
	}

	for _, app := range appList {
		_, err := exec.LookPath(app)
		if err != nil {
			t.Errorf("App %s was not installed: %v", app, err)
		}
	}
}