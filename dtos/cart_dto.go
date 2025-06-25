package dtos

type CartItemResponse struct {
	ItemID      uint    `json:"item_id"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
}

type GetCartResponse struct {
	CartItems  []CartItemResponse `json:"cart_items"`
	TotalPrice float64            `json:"total_price"`
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required,min=1"`
}

type AddToCartResponse struct {
	Message string `json:"message"`
}
