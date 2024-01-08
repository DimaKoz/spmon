package repository

import (
	"errors"
	"fmt"
)

// errNoHsRepo an error of HsStorage repository.
var errNoHsRepo = errors.New("couldn't find Handshake")

// errNoChkUnitRepo an error of ChkUnitStorage repository.
var errNoChkUnitRepo = errors.New("couldn't find the CheckUnit")

// repositoryError wraps error with msg and returns wrapped error.
func repositoryError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}
