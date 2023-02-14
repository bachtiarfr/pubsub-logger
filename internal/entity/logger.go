package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type TagLog string // internal, provider, activity

const (
	TagLogInternal = TagLog("internal")
	TagLogProvider = TagLog("provider")
	TagLogActivity = TagLog("activity")
)

type Logger struct {
	ID         		uuid.UUID       `json:"id"`          //uuid primarykey
	ServiceName		string	        `json:"service_name"`//varchar
	Request    		json.RawMessage `json:"request"`     // json
	RequestVendor   json.RawMessage `json:"request_vendor"`     // json
	Response   		json.RawMessage `json:"response"`    // json
	ResponseVendor  json.RawMessage `json:"response_vendor"`    // json
	URL        		string          `json:"url"`         // varchar
	UrlVendor       string          `json:"url_vendor"`         // varchar
	Method     		string          `json:"method"`      // varchar
	StatusCode 		int             `json:"status_code"` // int
	Event     		string 			`json:"event"`      // varchar
	RequestID 		string 			`json:"request_id"` // uuid
	Additional 		json.RawMessage `json:"additional"` // json
	Error 			json.RawMessage `json:"error"` //json
	Tag 			TagLog 			`json:"tag"`
	CreatedAt 		time.Time 		`json:"created_at"` // timestamp
}