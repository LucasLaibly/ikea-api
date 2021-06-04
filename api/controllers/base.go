package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver

	"github.com/LucasLaibly/ikea-api/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	//DBURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", DbUser, DbPassword, DbHost, DbPort, Dbdriver)

	// TODO:
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database. ", Dbdriver)
		log.Fatal(fmt.Sprintf("Error: %s", err))
	} else {
		fmt.Printf("Successfully connected to database.")
	}

	defer server.DB.Close()

	server.DB.Debug().AutoMigrate(&models.Customer{}, &models.Product{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	_ = godotenv.Load()

	fmt.Printf("On port %s", os.Getenv("DB_PORT"))

	log.Fatal(http.ListenAndServe(addr, server.Router))
}
