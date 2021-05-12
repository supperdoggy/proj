package itemsdata

import "github.com/supperdoggy/score/sctructs"

// CreateRequest - request struct for creating new item
type CreateRequest struct {
	Item sctructs.Item `json:"item"`
}

// CreateResponse - response struct for creating item
type CreateResponse struct {
	Item  sctructs.Item `json:"item"`
	Error string         `json:"error"`
}

// DeleteRequest - request struct for deleting item
// by id or by name
type DeleteRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// DeleteResponse - response struct for deleting item
type DeleteResponse struct {
	Error string         `json:"error"`
	Item  sctructs.Item `json:"item"`
}

// FindRequest - request struct for finding specific item in db
type FindRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// FindResponse - response struct for finding specific item in db
type FindResponse struct {
	Item  sctructs.Item `json:"item"`
	Error string         `json:"error"`
}
