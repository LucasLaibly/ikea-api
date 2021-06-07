package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID         string `gorm:"type:uuid;primary" json:"id"`
	Customer   Customer
	CustomerID Customer
	ProductID  Product
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

/*
Create uuid for cart?
*/
func (cart *Cart) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (cart *Cart) Prepare() {
	// todo?
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
