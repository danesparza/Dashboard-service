package api

import (
	"net/http"
)

var (
	// WsHub is the websocket hub
	WsHub = NewHub()
)

// ShowUI redirects to the '/ui/' virtual directory
func ShowUI(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/ui/", 301)
}

/*
// GetConfig gets a specfic config item based on application and config item name
func GetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := data.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	response, err := ds.Get(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	If we found an item, return it (otherwise, return an empty item):
	configItem := datastores.ConfigItem{}
	if response.Name != "" {
		configItem = response
		sendDataResponse(rw, "Config item found", configItem)
		return
	}

	sendDataResponse(rw, "No config item found with that application and name", configItem)
}

// SetConfig sets a specific config item
func SetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	response, err := ds.Set(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		WsHub.Broadcast <- []byte(getWSResponse("Updated", response))
		sendDataResponse(rw, "Config item updated", response)
	}
}

// RemoveConfig removes a specific config item
func RemoveConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	err = ds.Remove(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		WsHub.Broadcast <- []byte(getWSResponse("Removed", request))
		sendDataResponse(rw, "Config item removed", request)
	}
}

*/
