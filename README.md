# go-homebrew-apps-installer
This small script installs a list of homebrew packages(e.g. after formatting your MacOS device).  
  
You can add or remove casks/formulas from the list to suit your needs using flags or dialog menu through the terminal.  

See the list of packages in the config/config.go file.  

# Demo  
![demo](https://user-images.githubusercontent.com/95255319/232337780-acc96eff-ea49-49cd-be13-932c321b9bae.gif)

# Sample Usage
```ogi-app-arm64``` : Apple Silicon Chips   
```ogi-app-amd64``` : Intel and AMD Chips   
```./ogi-app-arm64  -a "adobe-acrobat-pro authy " -r "nordvpn zoom-us spotify"``` : Add adobe-acrobat-pro and authy to the list of packages to be installed and remove nordvpn, zoom-us and spotify from the list of packages to be installed.  
```./ogi-app-arm64  -all``` : Install all the packages directly without the dialogue menu.  

# Option 1 - Flags
The app can be used with flags to install the packages without the dialogue menu.  
```-a``` : Add package(s) to the list of packages to be installed.  
```-r``` : Remove package(s) from the list of packages to be installed.  
```-h``` : Show the help menu with all the available commands. 
```-all``` : Install all the packages directly without the dialogue menu.  

# Option 2 - Usage through the terminal dialogue menu
1. Download the build/ogi-app.zip and unzip it. (Alternatively,you can clone the repository and build the app yourself using the source code).  
2. Run the app by double clicking on it or through the terminal using the following command in the directory where the app is located:  ```./ogi-app-arm64``` or ```./ogi-app-amd64``` depending on your hardware architecture.
3. Follow the instructions on the terminal. Refer to the *Flags* section for more details on the usage with flags.  

# Running and building the app locally from the source code
```make run``` or ```go run cmd/main.go``` : Run the app locally.  
```make build``` or ```go build -o build/your-app-name cmd/main.go``` : Build a binary file of the app for ARM and Intel&AMD.  


