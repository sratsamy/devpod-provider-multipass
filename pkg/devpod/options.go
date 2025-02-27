package devpod

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MULTIPASS_PATH      = "MULTIPASS_PATH"
	MULTIPASS_IMAGE     = "MULTIPASS_IMAGE"
	MULTIPASS_CPUS      = "MULTIPASS_CPUS"
	MULTIPASS_DISK_SIZE = "MULTIPASS_DISK_SIZE"
	MULTIPASS_MEMORY    = "MULTIPASS_MEMORY"
	MULTIPASS_MOUNTS    = "MULTIPASS_MOUNTS"
	MULTIPASS_PORTS     = "MULTIPASS_PORTS"
)

type Options struct {
	Path     string // Path to multipass executable
	Image    string // --image arg passed into multipass launch command
	Cpus     int    // --cpus arg passed into multipass launch command
	DiskSize string // --disk arg passed into multipass launch command
	Memory   string // --memory arg passed into multipass launch command
	Mounts   string // host to multipass instance mounts
	Ports    string // New field
}

func OptsFromEnv() (*Options, error) {
	opts := &Options{}

	var err error

	opts.Path, err = fromEnvOrError(MULTIPASS_PATH)
	if err != nil {
		return nil, err
	}

	opts.Image, err = fromEnvOrError(MULTIPASS_IMAGE)
	if err != nil {
		return nil, err
	}

	cpus, err := fromEnvOrError(MULTIPASS_CPUS)
	if err != nil {
		return nil, err
	}

	opts.Cpus, err = strconv.Atoi(cpus)
	if err != nil {
		return nil, err
	}

	opts.DiskSize, err = fromEnvOrError(MULTIPASS_DISK_SIZE)
	if err != nil {
		return nil, err
	}

	opts.Memory, err = fromEnvOrError(MULTIPASS_MEMORY)
	if err != nil {
		return nil, err
	}

	opts.Mounts = os.Getenv(MULTIPASS_MOUNTS)

	opts.Ports = os.Getenv(MULTIPASS_PORTS)

	return opts, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf(
			"couldn't find option %s in environment, please make sure %s is defined",
			name,
			name,
		)
	}

	return val, nil
}
