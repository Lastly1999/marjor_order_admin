package request

type CreateOrderRequest struct {
	PhoneNum int    `json:"phoneNum"`
	City     string `json:"city"`
}
