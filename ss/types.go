package ss

import (
	"io"

	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/sei-protocol/sei-db/proto"
)

// StateStore is a versioned, embedded Key-Value Store,
// which allows efficient reads, writes, iteration over a specific version
type StateStore interface {
	Get(storeKey string, version uint64, key []byte) ([]byte, error)
	Has(storeKey string, version uint64, key []byte) (bool, error)
	Iterator(storeKey string, version uint64, start, end []byte) (types.Iterator, error)
	ReverseIterator(storeKey string, version uint64, start, end []byte) (types.Iterator, error)
	GetLatestVersion() (uint64, error)
	SetLatestVersion(version uint64) error

	// ApplyChangeset Persist the change set of a block,
	// the `changeSet` should be ordered by (storeKey, key),
	// the version should be latest version plus one.
	ApplyChangeset(version uint64, cs *proto.NamedChangeSet) error

	// Import the initial state of the store
	Import(version uint64, ch <-chan ImportEntry) error

	// Prune attempts to prune all versions up to and including the provided
	// version argument. The operation should be idempotent. An error should be
	// returned upon failure.
	Prune(version uint64) error

	// Closer releases associated resources. It should NOT be idempotent. It must
	// only be called once and any call after may panic.
	io.Closer
}

type ImportEntry struct {
	StoreKey string
	Key      []byte
	Value    []byte
}
