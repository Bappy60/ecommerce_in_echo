package models

import (
    "time"
)

type User struct {
    ID            uint64            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    Name          string            `json:"name" validate:"required,min=2,max=30"`
    Password      string            `json:"password"   validate:"required,min=6"`
    Email         string            `json:"email"      validate:"email,required"`
    Phone         string            `json:"phone"      validate:"required"`
    Token         string            `json:"token"`
    RefreshToken  string            `json:"refresh_token"`
    CreatedAt     time.Time         `json:"created_at"`
    UpdatedAt     time.Time         `json:"updated_at"`
    AddressID     uint64            `json:"address_id"`
    Address       Address           `json:"address"`
    UserCart      []SelectedProduct `json:"user_cart"`
    Orders        []Order           `json:"orders"`
}

type ProductCategory struct {
    ID          uint64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    CategoryName string `json:"category_name"`
    Description string  `json:"description"`
}


type Product struct {
    ID           uint64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    Name         string  `json:"name"`
    Category     ProductCategory `json:"category" gorm:"foreignkey:CategoryID"`
    CategoryID   uint64 `json:"category_id"`
    Price        uint64  `json:"price"`
}


type SelectedProduct struct {
    ID           uint64   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    ProductID    uint64   `json:"product_id"`
    Product      Product  `json:"product"`
    Price        float64  `json:"price"`
    OrderID      uint64   `json:"order_id"`
}

type Address struct {
    ID         uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    Country    string `json:"country"`
    City       string `json:"city"`
    PostalCode string `json:"postal_code"`
    Users      []User `json:"users"`
}

type Order struct {
    ID           uint64             `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    UserID       uint64             `json:"user_id"`
    User         User               `json:"user" gorm:"foreignkey:UserID"`
    OrderedAt    time.Time          `json:"ordered_at"`
    TotalPrice   float64            `json:"total_price"`
    Payment      Payment            `json:"payment"`
    SelectedProducts []SelectedProduct `json:"selected_products"`
}

type Payment struct {
    Digital bool   `json:"digital"`
    COD     bool   `json:"cod"`
}
