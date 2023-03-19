package entity

import "time"

type Source struct {
	Type        string                 `json:"type"`
	UserId      int64                  `json:"user_id,omitempty"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateUpdated time.Time              `json:"date_updated"`
}
