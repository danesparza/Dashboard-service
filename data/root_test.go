package data_test

import (
	"os"
	"testing"

	"github.com/danesparza/Dashboard-service/data"
)

//	Gets the database path for this environment:
func getTestFiles() string {
	systemdb := os.Getenv("DASH_TEST_ROOT")

	return systemdb
}

func TestRoot_GetTestDBPaths_Successful(t *testing.T) {

	systemdb := getTestFiles()

	if systemdb == "" {
		t.Fatal("The required DASH_TEST_ROOT environment variable is not set to the test database root path")
	}

	t.Logf("System db path: %s", systemdb)
}

func TestRoot_Databases_ShouldNotExistYet(t *testing.T) {
	//	Arrange
	systemdb := getTestFiles()

	//	Act

	//	Assert
	if _, err := os.Stat(systemdb); err == nil {
		t.Errorf("System database check failed: System db %s already exists, and shouldn't", systemdb)
	}
}

func TestRoot_GetKey_ReturnsCorrectKey(t *testing.T) {
	//	Arrange
	configKey := "unitestconfig1"
	expectedKey := "Config:unitestconfig1"

	//	Act
	actualKey := data.GetKey("Config", configKey)

	//	Assert
	if expectedKey != string(actualKey) {
		t.Errorf("GetKey failed:  Expected %s but got %s instead", expectedKey, actualKey)
	}
}
