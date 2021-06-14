package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID         string    `gorm:"type:uuid;primary" json:"id"`
	Customer   Customer  `json:"customer"`
	CustomerID string    `gorm:"not null;" json:"customerId"`
	Product    Product   `json:"product"`
	ProductID  string    `gorm:"not null;" json:"productId"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

/*
Create uuid for cart
*/
func (cart *Cart) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

/*
Prepare payload
*/
func (cart *Cart) Prepare() {
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
}

/*
Save Cart to Database
*/
func (cart *Cart) SaveCart(db *gorm.DB) (*Cart, error) {
	var err error
	err = db.Debug().Create(&cart).Error

	if err != nil {
		return &Cart{}, err
	}

	return cart, nil
}

/*
Get all items in cart for a customer
*/
func (c *Cart) GetAllItemsInCart(db *gorm.DB, customer_id string) error {
	return db.Raw("SELECT id, product_id FROM carts WHERE customer_id = ?", customer_id).Scan(c).Error
}
