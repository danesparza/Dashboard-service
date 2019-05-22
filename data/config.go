package data

import (
	"time"
)

// ConfigItem represents a single configuration item
type ConfigItem struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

/*
// AddConfig adds a config item to the system
func (store Manager) AddConfig(configName string, configDescription string) (ConfigItem, error) {
	//	Our return item
	retval := ConfigItem{}

	//	Our new config:
	config := ConfigItem{
		Name:        configName,
		Description: configDescription,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	//	First -- does the config item exist already?
	err := store.systemdb.View(func(txn *badger.Txn) error {
		_, err := txn.Get(GetKey("Config", config.Name))
		return err
	})

	//	If we didn't get an error, we have a problem:
	if err == nil {
		return retval, fmt.Errorf("Config item already exists")
	}

	//	Serialize to JSON format
	encoded, err := json.Marshal(role)
	if err != nil {
		return retval, fmt.Errorf("Problem serializing the data: %s", err)
	}

	//	Save it to the database:
	err = store.systemdb.Update(func(txn *badger.Txn) error {
		err := txn.Set(GetKey("Role", role.Name), encoded)
		return err
	})

	//	If there was an error saving the data, report it:
	if err != nil {
		return retval, fmt.Errorf("Problem saving the data: %s", err)
	}

	//	Set our retval:
	retval = role

	//	Return our data:
	return retval, nil
}

// GetRole gets a user from the system
func (store Manager) GetRole(context User, roleName string) (Role, error) {
	//	Our return item
	retval := Role{}

	//	Security check:  Are we authorized to perform this action?
	if !store.IsUserRequestAuthorized(context, sysreqGetRole) {
		return retval, fmt.Errorf("User %s is not authorized to perform the action", context.Name)
	}

	err := store.systemdb.View(func(txn *badger.Txn) error {
		item, err := txn.Get(GetKey("Role", roleName))
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

// GetAllRoles gets all roles in the system
func (store Manager) GetAllRoles(context User) ([]Role, error) {
	//	Our return item
	retval := []Role{}

	//	Security check:  Are we authorized to perform this action?
	if !store.IsUserRequestAuthorized(context, sysreqGetAllRoles) {
		return retval, fmt.Errorf("User %s is not authorized to perform the action", context.Name)
	}

	err := store.systemdb.View(func(txn *badger.Txn) error {

		//	Get an iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		//	Set our prefix
		prefix := GetKey("Role")

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
				item := Role{}

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
*/
