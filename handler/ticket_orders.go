// File: handler/ticket.go

package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// Replace with your actual package import path
)

func GetTicketDetails(c echo.Context) error {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ticket ID"})
	}

	// Assuming InitDB() initializes and returns a database connection with prepared queries
	ctx, queries := InitDB()
	details, err := queries.GetTicketDetails(ctx, int32(ticketID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve ticket details"})
	}

	return c.JSON(http.StatusOK, details)
}
