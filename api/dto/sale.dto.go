package dto

/*
sale and sale detail
define how transaction request structure object
cart
*/

type SaleCreate struct {
	Total     int64 `json:"total" binding:"required"`
	Cash      int64 `json:"Cash" binding:"required"`
	CashierID uint
}

type SaleDetailCreate struct {
	Qty    int16 `json:"quantity" binding:"required"`
	SaleID uint
	ItemId uint
}
