# hello-frog

## About this plugin
This plugin is a template and a functioning example for a basic JFrog CLI plugin. 
This README shows the expected structure of your plugin's README.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install JFrogTrafficMonitor 

Installing a specific version:

`$ jfrog plugin install JFrogTrafficMonitor@version

Uninstalling a plugin

`$ jfrog plugin uninstall JFrogTrafficMonitor `

Prerequisities:-
In the current directory, we need to have artifactory traffic logs usually in the format (artifactory-traffic.<epochtime>.log and artifactory-xray-traffic.<epochtime>.log)

## Usage
### Commands
* hello
    - Arguments:
        - addressee - The name of the person you would like to greet.
    - Flags:
        - shout: Makes output uppercase **[Default: false]**
        - repeat: Greets multiple times **[Default: 1]**
    - Example:
    ```
  $ jfrog hello-frog hello world --shout --repeat=2
  
  NEW GREETING: HELLO WORLD!
  NEW GREETING: HELLO WORLD!
  ```

### Environment variables
* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: New greeting: ]**

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).
