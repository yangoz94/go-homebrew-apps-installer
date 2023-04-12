package installers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
	"ogi/pkg/operations"
)

func InstallHomebrew() {
	_, err := exec.LookPath("brew")
	if err == nil {
		fmt.Println("Homebrew is already installed.")
		return
	}

	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Homebrew: %v", err)
	}
	fmt.Println("Homebrew has been installed.")
}

func InstallSelectedApps(appList *[]string) error {
	start := time.Now()
	var wg sync.WaitGroup
	for _, app := range *appList {
		wg.Add(1)
		go func(app string) {
			defer wg.Done()
			fmt.Printf("Installing %s...\n", app)
			if err := operations.RunCommand("env", "HOMEBREW_NO_AUTO_UPDATE=1", "brew", "install", app); err != nil {
				log.Fatal(err)
			}
			log.Printf("App %s installed successfully", app)
		}(app)
	}
	wg.Wait()
	fmt.Printf("All apps have been installed in %s\n", time.Since(start))
	return nil
}

