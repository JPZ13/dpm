# Docker Package Manager
[![CircleCI](https://circleci.com/gh/JPZ13/dpm.svg?style=shield)](https://circleci.com/gh/JPZ13/dpm)[![Go Report Card](https://goreportcard.com/badge/github.com/JPZ13/dpm)](https://goreportcard.com/report/github.com/JPZ13/dpm)


Tired of spending too much time configuring your machine? Me too.

DPM solves:
- Issues with different versions of node, python, golang etc
- System dependency conflicts
- Endlessly chaining things to your $PATH
- Random software updates that break everything

## Installation

New install instructions to come

## Usage

### Defining commands for your project

Add a file called `dpm.yml` to your project root defining the commands you want to use for development. For example:

```yaml
commands:
  go:
    image: golang:1.7.5
    entrypoints:
      - go

  python:
    image: python:alpine
    entrypoints:
      - python
      - pip
```

### Installing commands

Execute:
```
    dpm activate
```
DPM will then route any of the entrypoints in the current directory or child
directories to spin up a container using the specified image, attach it to a
named volume that holds the current directory and any child directories, run the
command in the volume-mounted container, and then stop and remove the container.
The end result is that your command line behaves as normal, but it runs commands
in containers


### Using project commands

Use commands as you normally would in the activated directory and any child
directories. The commands will only be run in containers in the activated
directories and not in outside directories.

### Deactivate
To go back to your normal system configuration,
```
  dpm deactivate
```

## Roadmap

See the `roadmap.md` file in the `/design` folder.

## Shoutout

Special thank you to [fermayo](https://github.com/fermayo) who did the initial [POC](https://github.com/fermayo/dpm). I heavily modified
the project later during Docker's Hack Week and continued working on it
subsequently. Also thank you to Krish for organizing Docker Hack Week, despite
leaving me off the ballot for voting. :grin:
