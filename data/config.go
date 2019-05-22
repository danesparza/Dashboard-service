package data

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/danesparza/badger"
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
	err = store.systemdb.Update(func(txn *badger.Txn) error {
		err := txn.Set(GetKey("Config", config.Name), encoded)
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
	err := store.systemdb.Update(func(txn *badger.Txn) error {
		err := txn.Delete(GetKey("Config", name))
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

	err := store.systemdb.View(func(txn *badger.Txn) error {
		item, err := txn.Get(GetKey("Config", name))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}

		if len(val) > 0 {
			//	Unmarshal data into our item
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

	err := store.systemdb.View(func(txn *badger.Txn) error {

		//	Get an iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		//	Set our prefix
		prefix := GetKey("Config")

		//	Iterate over our values:
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {

			//	Get the item
			item := it.Item()

			//	Get the item key
			// k := item.Key()

			//	Get the item value
			val, err := item.Value()
			if err != nil {
				return err
			}

			if len(val) > 0 {
				//	Create our item:
				item := ConfigItem{}

				//	Unmarshal data into our item
				if err := json.Unmarshal(val, &item); err != nil {
					return err
				}

				//	Add to the array of returned users:
				retval = append(retval, item)
			}
		}
		return nil
	})

	//	If there was an error, report it:
	if err != nil {
		return retval, fmt.Errorf("Problem getting the list of items: %s", err)
	}

	//	Return our data:
	return retval, nil
}
