# JFrogTrafficMonitor

## About this plugin
This plugin will show top20 uploads, top20 downloads and sum of the data usage. 

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install JFrogTrafficMonitor 

Installing a specific version:

`$ jfrog plugin install JFrogTrafficMonitor@version

Uninstalling a plugin

`$ jfrog plugin uninstall JFrogTrafficMonitor `

**Prerequisities**:-
In the current directory, we need to have artifactory traffic logs usually in the format (artifactory-traffic.epochtime.log and artifactory-xray-traffic.epochtime.log)

## Usage
### Commands
jfrog JFrogTrafficMonitor <command>


NAME:
   JFrogTrafficMonitor - Monitors all traffic of the repositories 

USAGE:
  jfrog JFrogTrafficMonitor [global options] command [command options] [arguments...]
   
VERSION:
   v0.0.1
   
COMMANDS:
   
    top20d, Top20D   This command results in the Top 20 downloads with no. of downloads, repository path and artifact size
   
    totalsum, TotalSum  This command results in the sum of data usage for the traffic logs present in the current directory.(total data tranfer)
   
    top20u, Top20U     This command results in the Top 20 uploads with no. of uploads, repository path and artifact size
   
     help, h          Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
   

## Release Notes
The release notes are available [here](RELEASE.md).
