package models

import (
	"time"
)

type (
	Project struct {
		Id     int    `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json: "id"`
		UserId int    `json: "user_id"`
		Name   string `json: "name"`
		Data   string `json: "data"`
		// WritersID
		// ReadersID
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
