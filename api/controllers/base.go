package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucasLaibly/ikea-api/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database.", Dbdriver)
		log.Fatal("Error: %s", err)
	} else {
		fmt.Printf("Successfully connected to database.")
	}

	server.DB.Debug().AutoMigrate(&models.Customer{}, &models.Product{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("On port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
