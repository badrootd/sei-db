package ss

import (
	"github.com/badrootd/sei-db/common/utils"
	"github.com/badrootd/sei-db/config"
	"github.com/badrootd/sei-db/ss/pebbledb"
	"github.com/badrootd/sei-db/ss/types"
)

func init() {
	initializer := func(dir string, configs config.StateStoreConfig) (types.StateStore, error) {
		dbHome := dir
		if configs.DBDirectory != "" {
			dbHome = configs.DBDirectory
		}
		return pebbledb.New(utils.GetStateStorePath(dbHome, configs.Backend), configs)
	}
	RegisterBackend(PebbleDBBackend, initializer)
}
