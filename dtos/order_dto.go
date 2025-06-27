package dtos

type CreateOrderRequest struct {
	Note string `json:"note"`
}

type CreateOrderItemDetail struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	Subtotal  float64 `json:"subtotal"`
}

type CreateOrderResponse struct {
	OrderID    uint                    `json:"order_id"`
	TotalPrice float64                 `json:"total_price"`
	Items      []CreateOrderItemDetail `json:"items"`
}
