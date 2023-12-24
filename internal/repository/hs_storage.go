package repository

import (
	"errors"
	"fmt"
	"sync"

	"spmon/internal/model"
)

var (
	hsStorageSync = &sync.Mutex{}
	hsStorage     = HsStorage{
		storage: make(map[string]model.Handshake, 0),
	}
)

// HsStorage represents storage of model.Handshake.
type HsStorage struct {
	storage map[string]model.Handshake
}

// errRepo an error of repository.
var errRepo = errors.New("couldn't find Handshake")

// repositoryError wraps error with msg and returns wrapped error.
func repositoryError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

// AddHs adds model.Handshake to 'hsStorage' storage.
func AddHs(key string, hs *model.Handshake) {
	hsStorageSync.Lock()
	defer hsStorageSync.Unlock()
	if hs == nil {
		delete(hsStorage.storage, key)
		return
	}
	addHsImpl(key, *hs)

	return
}

func addHsImpl(key string, hs model.Handshake) {
	hsStorage.storage[key] = hs

	return
}

// GetHs returns a *model.Handshake and nil error if found or nil and error.
func GetHs(key string) (*model.Handshake, error) {
	hsStorageSync.Lock()
	defer hsStorageSync.Unlock()

	if found, ok := hsStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errRepo, key)
}
