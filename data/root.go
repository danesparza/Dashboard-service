package data

import (
	"fmt"
	"strings"

	"github.com/danesparza/badger"
)

// Manager is the data manager
type Manager struct {
	systemdb *badger.DB
}

// WebSocketResponse represents a WebSocket event response
type WebSocketResponse struct {
	Type string     `json:"type"`
	Data ConfigItem `json:"data"`
}

// NewManager creates a new instance of a Manager and returns it
func NewManager(systemdbpath string) (*Manager, error) {
	retval := new(Manager)

	//	Open the systemDB
	sysopts := badger.DefaultOptions
	sysopts.Dir = systemdbpath
	sysopts.ValueDir = systemdbpath
	sysdb, err := badger.Open(sysopts)
	if err != nil {
		return retval, fmt.Errorf("Problem opening the systemDB: %s", err)
	}
	retval.systemdb = sysdb

	//	Return our Manager reference
	return retval, nil
}

// Close closes the data Manager
func (store Manager) Close() error {
	syserr := store.systemdb.Close()

	if syserr != nil {
		return fmt.Errorf("An error occurred closing the manager.  Syserr: %s", syserr)
	}

	return nil
}

// GetKey returns a key to be used in the storage system
func GetKey(entityType string, keyPart ...string) []byte {
	allparts := []string{}
	allparts = append(allparts, entityType)
	allparts = append(allparts, keyPart...)
	return []byte(strings.Join(allparts, ":"))
}
