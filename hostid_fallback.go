// +build !darwin,!linux,!freebsd,!windows

package nxid

import "errors"

func readPlatformMachineID() (string, error) {
	return "", errors.New("not implemented")
}
