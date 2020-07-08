package data

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tidwall/buntdb"
)

// ConfigItem represents a single configuration item
type ConfigItem struct {
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// SetConfig sets a config item to the system
func (store Manager) SetConfig(name string, value string) (ConfigItem, error) {
	//	Our return item
	retval := ConfigItem{}

	//	Our new config:
	config := ConfigItem{
		Name:    name,
		Value:   value,
		Created: time.Now(),
		Updated: time.Now(),
	}

	//	Serialize to JSON format
	encoded, err := json.Marshal(config)
	if err != nil {
		return retval, fmt.Errorf("Problem serializing the data: %s", err)
	}

	//	Save it to the database:
	err = store.systemdb.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(GetKey("Config", config.Name), string(encoded), nil)
		return err
	})

	//	If there was an error saving the data, report it:
	if err != nil {
		return retval, fmt.Errorf("Problem saving the data: %s", err)
	}

	//	Set our retval:
	retval = config

	//	Return our data:
	return retval, nil
}

// DeleteConfig removes a config item from the system
func (store Manager) DeleteConfig(name string) error {

	//	Save it to the database:
	err := store.systemdb.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(GetKey("Config", name))
		return err
	})

	//	If there was an error removing the data, report it:
	if err != nil {
		return fmt.Errorf("Problem removing the data: %s", err)
	}

	//	Return no error:
	return nil
}

// GetConfig gets a config item from the system
func (store Manager) GetConfig(name string) (ConfigItem, error) {
	//	Our return item
	retval := ConfigItem{}

	err := store.systemdb.View(func(tx *buntdb.Tx) error {
		item, err := tx.Get(GetKey("Config", name))
		if err != nil {
			return err
		}

		if len(item) > 0 {
			//	Unmarshal data into our item
			val := []byte(item)
			if err := json.Unmarshal(val, &retval); err != nil {
				return err
			}
		}

		return nil
	})

	//	If there was an error, report it:
	if err != nil {
		return retval, fmt.Errorf("Problem getting the data: %s", err)
	}

	//	Return our data:
	return retval, nil
}

// GetAllConfig gets all config items in the system
func (store Manager) GetAllConfig() ([]ConfigItem, error) {
	//	Our return item
	retval := []ConfigItem{}

	//	Set our prefix
	prefix := GetKey("Config")

	err := store.systemdb.View(func(tx *buntdb.Tx) error {

		tx.Descend(prefix, func(key, val string) bool {

			if len(val) > 0 {
				//	Create our item:
				item := ConfigItem{}

				//	Unmarshal data into our item
				bval := []byte(val)
				if err := json.Unmarshal(bval, &item); err != nil {
					return false
				}

				//	Add to the array of returned users:
				retval = append(retval, item)
			}

			return true
		})
		return nil
	})

	//	If there was an error, report it:
	if err != nil {
		return retval, fmt.Errorf("Problem getting the list of items: %s", err)
	}

	//	Return our data:
	return retval, nil
}
