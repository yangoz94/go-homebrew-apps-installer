package installers

import (
	"fmt"
	"log"
	"ogi/pkg/operations"
	"os"
	"os/exec"
	"time"
)

func InstallHomebrew() {
	_, err := exec.LookPath("brew")
	if err == nil {
		fmt.Println("Homebrew is already installed.")
		return
	}

	cmd := exec.Command("/bin/zsh", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Homebrew: %v", err)
	}
	fmt.Println("Homebrew has been installed.")
}

func InstallSelectedApps(appList *[]string) error {
	start := time.Now()
	for _, app := range *appList {
		fmt.Printf("Installing %s...\n", app)
		if err := operations.RunCommand("env", "HOMEBREW_NO_AUTO_UPDATE=1", "brew", "install", app); err != nil {
			log.Fatal(err)
		}
		log.Printf("App %s installed successfully", app)
	}
	fmt.Printf("All apps have been installed in %s\n", time.Since(start))
	return nil
}
