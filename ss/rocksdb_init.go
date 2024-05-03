//go:build rocksdbBackend
// +build rocksdbBackend

package ss

import (
	"github.com/badrootd/sei-db/common/utils"
	"github.com/badrootd/sei-db/config"
	"github.com/badrootd/sei-db/ss/rocksdb"
	"github.com/badrootd/sei-db/ss/types"
)

func init() {
	initializer := func(dir string, configs config.StateStoreConfig) (types.StateStore, error) {
		dbHome := dir
		if configs.DBDirectory != "" {
			dbHome = configs.DBDirectory
		}
		return rocksdb.New(utils.GetStateStorePath(dbHome, configs.Backend), configs)
	}
	RegisterBackend(RocksDBBackend, initializer)
}
