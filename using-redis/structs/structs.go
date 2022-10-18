package structs

import "github.com/gofrs/uuid"

type Users struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	CreatedTime string    `json:"created_time"`
	UpdatedTime string    `json:"updated_time"`
	Source      string    `json:"source"`
}

type Error struct {
	Message string `json:"error"`
}
