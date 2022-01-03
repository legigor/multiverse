package controller

import (
	"os"
	"path/filepath"
)

func getKubeConfigPath() (string, error) {
	path, ok := os.LookupEnv("KUBECONFIG")
	if ok {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = filepath.Join(homeDir, ".kube/config")

	return path, nil
}
