package model

type Order struct {
	Base
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
