# Airbyte CLI

This is the next generation Go based Airbyte CLI.

## Contents
- [Setup](#setup)
- [Configuration as YAML](#configuration-as-yaml)
- [Usage](#usage)
  - [get](#get)
  - [search](#search)
  - [check](#check)
  - [logs](#logs)
  - [export](#export)

##  Setup
1. Clone this repo to your working environment  
   `git clone https://github.com/harshithmullapudi/airbyte-cli.git`
2. Install go (if not installed already).  
   `brew install go` (for mac users)
3. Jump to the airbyte-cli directory and run
   `go install .`
4. You're all set to fly

##  Configuration as YAML
Before we interact with Airbyte API, we need to set Airbyte API URL and workspace, for which you can use the following command
```bash
$ airbyte set-config
```

##  Usage
An Airbyte CLI command has the following structure:
```bash
$ airbyte <command> <subcommand> [options and parameters]
```
```
<command>
          get - Get configuration of Sources/Destinations/Connections
          search - Search in sources
          check - Check connection to Source/Destination
          logs - Fetch logs for a job
          set-config - Set your airbyte url and workspaceID
          help - Help about any command
```
To view help documentation, use one of the following:
```bash
$ airbyte -h
$ airbyte <command> -h
$ airbyte <command> <subcommand> -h
```
## get
Return all
   - sources (`/v1/sources/list`)
   - destinations (`/v1/destinations/list`)
   - connections (`/v1/web_backend/connections/list`)

You can use page(p) and offset(o) to fetch sources respectively. 

To list sources, the command would be:
```bash
$ airbyte get sources
```
To get a source, the command would be:
```bash
$ airbyte get source [sourceId]
```
To get jobs for a connection, the command would be:
```bash
$ airbyte get jobs [connectionId]
```

## search
   - sources
   - connections

To search in sources, the command would be:
```bash
$ airbyte search sources [string]
```

## check
Validate
   - source (`/v1/sources/check_connection`)
   - destination (`/v1/destinations/check_connection`)

To validate a source, the command would be:
```bash
$ airbyte check source [sourceId]
```

## logs
Get logs for a job
```bash
$ airbyte logs [jobId]
```

## export
Export sources, connections and destinations to a target folder
```bash
$ airbyte export -t [absolute path to target folder]
```

## create

You can create and validation sources through a yaml file. 

1. Create a folder in your home example: /Users/name/home/load
2. Create a yaml file SOURCE_CONNECTION.yaml


Validate sources before creation
```bash
$ airbyte create -f [path to config folder]
```

You can pass -c as a flag to create the sources which will validate the source and skip it either if validation failes or if it finds sourceId and 
a source for that respective Id.

```bash
$ airbyte create -f [path to config folder]
```