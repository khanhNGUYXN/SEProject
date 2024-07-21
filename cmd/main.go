package main

import (
	"log"

	"example.com/test/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// General and news-related routes
	e.GET("/", handler.GetIndex)
	e.GET("/home", handler.GetHome)
	e.GET("/home/news", handler.GetNews)
	e.GET("/home/movies", handler.GetMovieSelection)

	// Movie-specific routes
	e.GET("/movie/:id/booking", handler.GetBookingView) // For viewing booking options for a specific movie
	e.GET("/movie/:id/info", handler.GetMovieInfo)      // For getting detailed information about a movie

	// Login and authentication routes
	e.GET("/login", handler.GetLoginPage)
	e.POST("/login/auth", handler.CheckLogin) // Endpoint to handle login authentication

	// CRUD operations for movies (Assuming you want these routes)
	e.POST("/movies", handler.CreateMovie)       // Create a new movie
	e.PUT("/movies/:id", handler.UpdateMovie)    // Update an existing movie
	e.DELETE("/movies/:id", handler.DeleteMovie) // Delete a movie
	e.GET("/tickets/:id", handler.GetTicketDetails)
	e.GET("/notify", handler.NotifyCustomer)

	e.POST("/coupons", handler.AddCoupon)
	e.GET("/coupons", handler.ListCoupons)

	// Static files handler
	e.Static("/static", "static")

	// Start the server
	log.Fatal(e.Start(":8080"))
}
