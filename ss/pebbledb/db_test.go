package pebbledb

import (
	"testing"

	"github.com/badrootd/sei-db/config"
	sstest "github.com/badrootd/sei-db/ss/test"
	"github.com/badrootd/sei-db/ss/types"
	"github.com/stretchr/testify/suite"
)

func TestStorageTestSuite(t *testing.T) {
	s := &sstest.StorageTestSuite{
		NewDB: func(dir string) (types.StateStore, error) {
			return New(dir, config.DefaultStateStoreConfig())
		},
		EmptyBatchSize: 12,
	}

	suite.Run(t, s)
}
