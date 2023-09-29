package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderID string
type ProductID string
type VendorID string

// Order define the app order data
type Order struct {
	ID        OrderID   `json:"ID"`
	ProductId ProductID `json:"ProductID"`
	CreatedAt time.Time `json:"CreatedAt"`
	FullName  string    `json:"FullName"`
	OrderDate time.Time `json:"OrderDate"`
}

// Product define the app product data
type Product struct {
	ID          ProductID       `json:"ID"`
	VendorId    VendorID        `json:"VendorID"`
	CreatedAt   time.Time       `json:"CreatedAt"`
	Title       string          `json:"Title"`
	ListingType string          `json:"Listing_Type"`
	Price       decimal.Decimal `json:"Price"`
}

// Vendor define the app vendor data
type Vendor struct {
	ID        VendorID  `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Name      string    `json:"Name"`
}
