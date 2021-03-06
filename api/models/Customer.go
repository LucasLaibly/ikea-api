package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	ID        string    `gorm:"type:uuid;primary" json:"id"`
	Name      string    `gorm:"size:30;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

/*
Gorm Hook, BeforeCreate Customer, generate a new uuid
*/
func (customer *Customer) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

/*
Prepare model
*/
func (customer *Customer) Prepare() {
	customer.Name = html.EscapeString(strings.TrimSpace(customer.Name))
	customer.Email = html.EscapeString(strings.TrimSpace(customer.Email))
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
}

/*
Validate data provided
*/
func (customer *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	default:
		if customer.Name == "" {
			return errors.New("Name Required.")
		}
		if customer.Email == "" {
			return errors.New("Email Required.")
		}
		return nil
	}
}

/*
Notes for the future:
first  parameter set is method ownership  (acting on Customer)
second parameter set is the arugments for the function
third  parameter set is a multi-value return from the function. Both a Customer and an error in this case
*/
func (customer *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {
	var err error
	err = db.Debug().Create(&customer).Error

	if err != nil {
		return &Customer{}, err
	}

	return customer, nil
}

/*
Find user by their ID
*/
func (customer *Customer) FindCustomerByID(db *gorm.DB, uid string) (*Customer, error) {
	var err error

	err = db.Debug().Model(Customer{}).Where("id = ?", uid).Take(&customer).Error

	if err != nil {
		return &Customer{}, err
	}

	return customer, nil
}
