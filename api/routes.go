package api

import "net/http"

var (
	// WsHub is the websocket hub
	WsHub = NewHub()
)

// ShowUI redirects to the '/ui/' virtual directory
func ShowUI(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/ui/", 301)
}
