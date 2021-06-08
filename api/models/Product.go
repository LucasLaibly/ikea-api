package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          string    `gorm:"type:uuid;primary" json:"id"`
	Name        string    `gorm:"size:30;not null;unique" json:"name"`
	Description string    `gorm:"size:200;not null;unique" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deletedAt"`
}

/*
Gorm Hook, BeforeCreate Product, generate new uuid
*/
func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

/*
Prepare model
*/
func (product *Product) Prepare() {
	product.Name = html.EscapeString(strings.TrimSpace(product.Name))
	product.Description = html.EscapeString(strings.TrimSpace(product.Description))
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.DeletedAt = time.Now()
}

/*
Validate
*/
func (product *Product) Validate(action string) error {
	switch strings.ToLower(action) {
	default:
		if product.Name == "" {
			return errors.New("Name Required.")
		}
		if product.Description == "" {
			return errors.New("Description Required.")
		}
		return nil
	}
}

/*
Save record
*/
func (product *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Create(&product).Error

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}

/*
Find product by it's ID
*/
func (product *Product) FindProductByID(db *gorm.DB, uid string) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&product).Error

	if err != nil {
		return &Product{}, err
	}

	return product, nil
}
