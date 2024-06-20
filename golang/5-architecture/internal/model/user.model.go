package model

type User struct {
	Base
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
