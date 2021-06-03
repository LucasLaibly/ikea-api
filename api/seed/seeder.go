package seed

import (
	"log"

	"github.com/LucasLaibly/ikea-api/api/models"
	"github.com/jinzhu/gorm"
)

var customers = []models.Customer{
	{
		Name:  "Emma Lejon",
		Email: "emma.lejon@email.com",
	},
	{
		Name:  "Lucas Laibly",
		Email: "lucas.laibly@gmail.com",
	},
}

var products = []models.Product{
	{
		Name:        "Table",
		Description: "Gather the family.",
	},
	{
		Name:        "Chair",
		Description: "Something I could use right about now.",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Customer{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Customer{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range customers {
		err = db.Debug().Model(&models.Customer{}).Create(&customers[i]).Error
		if err != nil {
			log.Fatalf("cannot see customers table: %v", err)
		}

		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products tables: %v", err)
		}
	}
}
