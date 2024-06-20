package model

type Image struct {
	Base
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
