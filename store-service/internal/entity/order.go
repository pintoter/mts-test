package entity

import "time"

type Order struct {
	ID        int64     `json:"id,omitempty"`
	UserId    int64     `json:"user_id"`
	ItemId    int64     `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
}
