[![CircleCI](https://circleci.com/gh/jsenon/kubextractor.svg?style=svg&circle-token=a8df7aa6e5450279a39e313b309dfd869472c9a5)](https://circleci.com/gh/jsenon/kubextractor)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/kubextractor)](https://goreportcard.com/report/github.com/jsenon/kubextractor)

# KubExtract

This tool could be used to extract specific k8s configuration context from global config file

## Build

```
go build kubextractor
```

### Use package

```
go get -u github.com/jsenon/kubextractor
```

### Prerequisite

You need to have kubectl command line properly installed and set to your PATH

### Exec

Download kubextractor executable file in bin/

### Command line

```
Extract kubernetes context ie. configuration user and endpoint.
				  Complete documentation is available at https://github.com/jsenon/kubextractor
				  After export Use kubectl config use-context YOURCONTEXT --kubeconfig output.json to use it

Usage:
  kubextractor [flags]

Flags:
  -c, --config string       k8s config file default ($HOME/.kube/config)
  -j, --configjson string   k8s config file JSON default ($HOME/.kube/config.json)
  -e, --context string      MANDATORY: Name of  context to extract
  -h, --help                help for kubextractor
  -o, --output string       Name of output file
  ```
  

### TIPS

Use Help if need

```
kubextractor -h
```

