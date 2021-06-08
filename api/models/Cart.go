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
Get all items in a customer's cart

func (c *Cart) GetAllItemsInCart(db *gorm.DB, uuid string) (*[]Cart, error) {
	var err error

	cart := []Cart{}

	err = db.Debug().Model(&Cart{}).Limit(100).Find(&c).Error

	if err != nil {
		return &[]Cart{}, err
	}

	if len(cart) > 0 {
		for i := range cart {
			err := db.Debug().Model(&Cart{}).Where("id = ?", cart[i].CustomerID).Take(&cart[i].Customer).Error
			if err != nil {
				return &[]Cart{}, err
			}
		}
	}

	return &cart, nil
}
*/
