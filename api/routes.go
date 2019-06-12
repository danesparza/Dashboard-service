package api

import (
	"encoding/json"
	"net/http"

	"github.com/danesparza/Dashboard-service/data"
)

var (
	// WsHub is the websocket hub
	WsHub = NewHub()
)

// ShowUI redirects to the '/ui/' virtual directory
func ShowUI(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/ui/", 301)
}

// GetConfig gets a specfic config item based on config item name
func (service Service) GetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := data.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Send the request to the datastore and get a response:
	response, err := service.DB.GetConfig(request.Name)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	If we found an item, return it (otherwise, return an empty item):
	configItem := data.ConfigItem{}
	if response.Name != "" {
		configItem = response
		sendDataResponse(rw, "Config item found", configItem)
		return
	}

	sendDataResponse(rw, "No config item found with that name", configItem)
}

// GetAllConfig gets all config items
func (service Service) GetAllConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := data.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Send the request to the datastore and get a response:
	response, err := service.DB.GetAllConfig()
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	If we found an item, return it (otherwise, return an empty item):
	configItems := []data.ConfigItem{}
	if len(response) == 0 {
		configItems = response
		sendDataResponse(rw, "Config items found", configItems)
		return
	}

	sendDataResponse(rw, "No config items found with that name", configItems)
}

// SetConfig sets a specific config item
func (service Service) SetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := data.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Send the request to the datastore and get a response:
	response, err := service.DB.SetConfig(request.Name, request.Value)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		WsHub.Broadcast <- []byte(getWSResponse("Updated", response))
		sendDataResponse(rw, "Config item updated", response)
	}
}

// RemoveConfig removes a specific config item
func (service Service) RemoveConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := data.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Send the request to the datastore and get a response:
	err = service.DB.DeleteConfig(request.Name)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		WsHub.Broadcast <- []byte(getWSResponse("Removed", request))
		sendDataResponse(rw, "Config item removed", request)
	}
}
