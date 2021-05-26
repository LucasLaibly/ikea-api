package models

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	ID        string    `pgsql:"type:uuid;primary;default:uuid_generate_v4()"`
	Name      string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (customer *Customer) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (customer *Customer) Prepare() {
	customer.Name = html.EscapeString(strings.TrimSpace(customer.Name))
	customer.Email = html.EscapeString(strings.TrimSpace(customer.Email))
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
}
