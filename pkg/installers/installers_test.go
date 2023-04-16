package installers

import (
	"log"
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
}

func TestInstallSelectedApps(t *testing.T) {
	//set the timeout to 3 minutes for testability 
	// (brew install can take a long time on slower machines)
	appList := []string{"lynx", "gray"}
	err := InstallSelectedApps(&appList)
	if err != nil {
		t.Errorf("Error installing apps: %v", err)
	}

	//uninstall apps
	for _, app := range appList {
		_, err := exec.Command("brew", "uninstall", app).Output()
		log.Printf("Uninstalling %s...", app)
		if err != nil {
			t.Errorf("Error uninstalling app %s: %v", app, err)
		}
	}
	log.Println("All apps uninstalled successfully")
}

