package controllers

import "github.com/LucasLaibly/ikea-api/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Carts
	s.Router.HandleFunc("/customers/{id}/cart", middlewares.SetMiddlewareJSON(s.GetAllCartItems)).Methods("GET")
	s.Router.HandleFunc("/cart", middlewares.SetMiddlewareJSON(s.CreateCart)).Methods("POST")

	// Customers
	s.Router.HandleFunc("/customers/{id}", middlewares.SetMiddlewareJSON(s.FindCustomer)).Methods("GET")
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")

	// Products
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.FindProductByID)).Methods("GET")
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
}
