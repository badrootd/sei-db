package ss

import (
	"github.com/badrootd/sei-db/common/utils"
	"github.com/badrootd/sei-db/config"
	"github.com/badrootd/sei-db/ss/pebbledb"
	"github.com/badrootd/sei-db/ss/types"
)

func init() {
	initializer := func(dataDir string, configs config.StateStoreConfig) (types.StateStore, error) {
		dbDir := dataDir
		if configs.DBDirectory != "" {
			dbDir = configs.DBDirectory
		}
		return pebbledb.New(utils.GetStateStorePath(dbDir, configs.Backend), configs)
	}
	RegisterBackend(PebbleDBBackend, initializer)
}
