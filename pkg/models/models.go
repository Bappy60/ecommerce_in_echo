package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID        uint64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	Email     string  `json:"email" gorm:"unique; not null"`
	AddressID uint64  `json:"address_id"`
	Address   Address `json:"address"`
	Carts     []Cart  `json:"carts"`
	// Orders    []Order `json:"orders"`
	HasRole   string  `json:"roles" gorm:"default:'user'"`
}


type ProductCategory struct {
	ID           uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

type Product struct {
	ID         uint64          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name       string          `json:"name"`
	Category   ProductCategory `json:"category" gorm:"foreignkey:CategoryID"`
	CategoryID uint64          `json:"category_id"`
	Price      float64         `json:"price"`
}

type Cart struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint64     `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CartItems []CartItem `json:"cart_items" gorm:"foreignkey:CartID"`
	User      User       `gorm:"foreignkey:UserID"`
}

type CartItem struct {
	ID             uint64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CartID         uint64  `gorm:"not null" json:"cart_id"`
	ProductID      uint64  `gorm:"not null" json:"product_id"`
	Quantity       uint64  `gorm:"not null" json:"quantity"`
	TotalItemPrice float64 `gorm:"not null" json:"total_item_price"`
	Product        Product `gorm:"foreignkey:ProductID" json:"product"`
}

// type Order struct {
// 	ID            uint64      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
// 	UserID        uint64      `gorm:"not null" json:"user_id"`
// 	AddressID     uint64      `gorm:"not null" json:"address_id"`
// 	TotalBill     float64     `gorm:"not null" json:"total_price"`
// 	OrderStatus   string      `gorm:"not null" json:"order_status"`
// 	PaymentStatus bool        `gorm:"not null" json:"payment_status"`
// 	OrderedAt     time.Time   `json:"ordered_at"`
// 	Address       Address     `gorm:"foreignkey:AddressID" json:"address"`
// 	OrderItems    []OrderItem `json:"order_items"`
// }

// type OrderItem struct {
// 	ID        uint64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
// 	OrderID   uint64  `gorm:"not null" json:"order_id"`
// 	ProductID uint64  `gorm:"not null" json:"product_id"`
// 	Quantity  uint32  `gorm:"not null" json:"quantity"`
// 	Price     float64 `gorm:"not null" json:"price"`
// 	Product   Product `json:"product"`
// }

type Address struct {
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func (pc ProductCategory) Validate() error {
	return validation.ValidateStruct(&pc,
		validation.Field(&pc.CategoryName, validation.Required),
		validation.Field(&pc.Description, validation.Required),
	)
}
