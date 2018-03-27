package model

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

// Bookmark type represents the metadata of a bookmark.
type Bookmark struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Priority    int `json:"priority"` // Priority (1 -5)
	CreatedBy   string `json:"createdby"`
	CreatedOn   time.Time `json:"createdon,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
