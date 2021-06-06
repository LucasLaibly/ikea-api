package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver

	"github.com/LucasLaibly/ikea-api/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database. ", Dbdriver)
		log.Fatal(fmt.Sprintf("Error: %s", err))
	} else {
		fmt.Printf("Successfully connected to database.")
	}

	//defer server.DB.Close()

	// auto migrations
	server.DB.Debug().AutoMigrate(&models.Customer{})
	server.DB.Debug().AutoMigrate(&models.Product{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("On port %s", addr)

	log.Fatal(http.ListenAndServe(addr, server.Router))
}
