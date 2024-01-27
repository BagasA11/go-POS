package dto

// retail and cashier
/*Retail registration structure*/
type RetailRegister struct {
	Email    string `json:"email" binding:"required,min=5"`
	Contact  string `json:"contact" binding:"required,min=7"`
	Password string `json:"password" binding:"required,min=8"`
}

type RetailLogin struct {
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type RetailUpdate struct {
	Email    string `json:"email" binding:"required,min=5"`
	Contact  string `json:"contact" binding:"required,min=7"`
	Password string `json:"password" binding:"required,min=8"`
}

type RetailNewPassword struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	Password    string `json:"password" binding:"required,min=8"`
}

type RetailDelete struct {
	Password string `json:"password" binding:"required,min=8"`
}

type CashierCreate struct {
	Username string `json:"username" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type CashierLogin struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

type CashierUpdate struct {
	Username string `json:"username" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type CashierNewPassword struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	Password    string `json:"password" binding:"required,min=8"`
}

type CashierDelete struct {
	Password string `json:"password" binding:"required,min=8"`
}
