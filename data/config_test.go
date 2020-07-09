package data_test

import (
	"os"
	"testing"

	"github.com/danesparza/Dashboard-service/data"
)

func TestConfig_SetConfig_ValidConfig_Successful(t *testing.T) {

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

	if response.Value != "Value1" {
		t.Errorf("SetConfig - Response should match set value, but got: %s", response.Value)
	}

}

func TestConfig_GetConfig_ValidConfig_Successful(t *testing.T) {

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
	configitem, err := db.SetConfig("Name1", "Value1")
	if err != nil {
		t.Errorf("SetConfig - Should execute without error, but got: %s", err)
	}

	response, err := db.GetConfig("Name1")

	//	Assert
	if err != nil {
		t.Errorf("GetConfig - Should execute without error, but got: %s", err)
	}

	if configitem.Value != response.Value {
		t.Errorf("GetConfig - Should match set value, but got: %s", response.Value)
	}

}

func TestConfig_GetAllConfig_ValidConfig_Successful(t *testing.T) {

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
	configitem1, err := db.SetConfig("Name1", "Value1")
	if err != nil {
		t.Errorf("GetAllConfig - Should add item without error, but got: %s", err)
	}
	t.Logf("Added %+v", configitem1)

	configitem2, err := db.SetConfig("Name2", "Value2")
	if err != nil {
		t.Errorf("GetAllConfig - Should add item without error, but got: %s", err)
	}
	t.Logf("Added %+v", configitem2)

	response, err := db.GetAllConfig()

	//	Assert
	if err != nil {
		t.Errorf("GetAllConfig - Should execute without error, but got: %s", err)
	}

	if len(response) < 2 {
		t.Errorf("GetAllConfig - Should get 2 items, but got: %v", len(response))
	}

	t.Logf("Get all configs: %+v", response)

}
