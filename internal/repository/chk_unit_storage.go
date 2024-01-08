package repository

import (
	"sync"

	"github.com/DimaKoz/spmon/internal/model"
	"golang.org/x/exp/maps"
)

var (
	chkUnitStorageSync = &sync.Mutex{}
	chkUnitStorage     = ChkUnitStorage{
		storage: make(map[string]model.CheckUnit, 0),
	}
)

// ChkUnitStorage represents a storage of model.CheckUnit.
type ChkUnitStorage struct {
	storage map[string]model.CheckUnit
}

// AddCheckUnit adds model.CheckUnit to 'repository.ChkUnitStorage' storage.
func AddCheckUnit(key string, chkUnit *model.CheckUnit) {
	chkUnitStorageSync.Lock()
	defer chkUnitStorageSync.Unlock()
	if chkUnit == nil {
		delete(chkUnitStorage.storage, key)

		return
	}
	addChkUnitImpl(key, *chkUnit)
}

func addChkUnitImpl(key string, chkUnit model.CheckUnit) {
	chkUnitStorage.storage[key] = chkUnit
}

// GetCheckUnit returns a *model.CheckUnit and nil error if found
// or nil and wrapped repository.errNoChkUnitRepo.
func GetCheckUnit(key string) (*model.CheckUnit, error) {
	chkUnitStorageSync.Lock()
	defer chkUnitStorageSync.Unlock()

	if found, ok := chkUnitStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errNoChkUnitRepo, key)
}

// ClearUnitStorage clear the storage.
func ClearUnitStorage() {
	chkUnitStorageSync.Lock()
	defer chkUnitStorageSync.Unlock()
	maps.Clear(chkUnitStorage.storage)
}

// GetAllChkUnits gets all the items from the storage.
func GetAllChkUnits() []model.CheckUnit {
	chkUnitStorageSync.Lock()
	defer chkUnitStorageSync.Unlock()

	return maps.Values(chkUnitStorage.storage)
}
