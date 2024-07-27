package model

type Transaction struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	UserID    int     `json:"user_id"`
	CartID    int     `json:"cart_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"` // e.g., "pending", "completed", "failed"
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type PaymentMethod struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UserID  int    `json:"user_id"`
	Method  string `json:"method"` // e.g., "credit_card", "paypal"
	Details string `json:"details"`
}
