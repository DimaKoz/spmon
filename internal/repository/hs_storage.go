package repository

import (
	"sync"

	"github.com/DimaKoz/spmon/internal/model"
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

// AddHs adds model.Handshake to 'hsStorage' storage.
func AddHs(key string, handshake *model.Handshake) {
	hsStorageSync.Lock()
	defer hsStorageSync.Unlock()
	if handshake == nil {
		delete(hsStorage.storage, key)

		return
	}
	addHsImpl(key, *handshake)
}

func addHsImpl(key string, hs model.Handshake) {
	hsStorage.storage[key] = hs
}

// GetHs returns a *model.Handshake and nil error if found or nil and error.
func GetHs(key string) (*model.Handshake, error) {
	hsStorageSync.Lock()
	defer hsStorageSync.Unlock()

	if found, ok := hsStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errNoHsRepo, key)
}
