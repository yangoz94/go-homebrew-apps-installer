# go-homebrew-apps-installer
This small script installs a list of my favorite homebrew packages at once after formatting your MacOS device.  
You can add or remove casks/formulas from the list to suit your needs using flags or dialog menu through the terminal.  

See the list of packages in the config/config.go file.  

# FLAGS
The app can be used with flags to install the packages without the dialogue menu.  
```-a``` : Add package(s) to the list of packages to be installed.  
```-r``` : Remove package(s) from the list of packages to be installed.  
```-h``` : Show the help menu.  
```-all``` : Install all the packages directly without the dialogue menu.  

# USAGE THROUGH TERMINAL DIALOGUE
1. Download the build/ogi-app file from the repository.(Alternatively,you can clone the repository and build the app yourself using the source code).  
2. Run the app by double clicking on it or through the terminal using the following command in the directory where the app is located:  ```./ogi-app```.  
3. Follow the instructions on the terminal. Refer to the *Flags* section for more details on the usage with flags.  



