package data_test

import (
	"os"
	"testing"

	"github.com/danesparza/Dashboard-service/data"
)

func TestConfig_AddConfig_ValidConfig_Successful(t *testing.T) {

	//	Arrange
	systemdb := getTestFiles()
	db, err := data.NewManager(systemdb)
	if err != nil {
		t.Errorf("NewManager failed: %s", err)
	}
	defer func() {
		db.Close()
		os.RemoveAll(systemdb)
	}()

	//	Act
	response, err := db.SetConfig("Name1", "Value1")

	//	Assert
	if err != nil {
		t.Errorf("SetConfig - Should execute without error, but got: %s", err)
	}

	t.Logf("Set config: %+v", response)

}
