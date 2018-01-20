# KubExtract

This tool could be used to extract specific k8s configuration context from global config file

## Build

```
go build kubextractor
```

### Prerequisite

You need to have kubectl command line properly installed and set to your PATH

### Exec

Download kubextractor executable file in bin/

### Command line

```
Extract kubernetes context ie. configuration user and endpoint.
				  Complete documentation is available at https://github.com/jsenon/kubextractor
				  Use kubectl config view -o json --raw --kubeconfig YOURCONFIG > output.json to generate json
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

