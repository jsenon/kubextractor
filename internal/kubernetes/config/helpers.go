package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// LoadConfig return a go struct representing the kubectl config file respecting kubernetes convention.
// Largely inspired from github.com/kubernetes/kubectl/pkg/pluginutils/plugin_client.go
func LoadConfig() (*clientcmdapi.Config, error) {
	// resolve kubeconfig location, prioritizing the --config global flag,
	// then the value of the KUBECONFIG env var (if any), and defaulting
	// to ~/.kube/config as a last resort.
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	kubeconfig := filepath.Join(home, ".kube", "config")

	kubeconfigEnv := os.Getenv("KUBECONFIG")
	if len(kubeconfigEnv) > 0 {
		kubeconfig = kubeconfigEnv
	}

	if len(kubeconfig) == 0 {
		return nil, fmt.Errorf("error initializing config. the KUBECONFIG environment variable must be defined.")
	}

	config, err := configFromPath(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error obtaining kubectl config: %v", err)
	}

	return config, nil
}

func configFromPath(path string) (*clientcmdapi.Config, error) {
	rules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: path}
	cfg, err := rules.Load()
	if err != nil {
		return nil, fmt.Errorf("the provided credentials %q could not be loaded: %v", path, err)
	}

	return cfg, nil
}
