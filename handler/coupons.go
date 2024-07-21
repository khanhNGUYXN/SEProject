// File: handler/coupons.go

package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// Adjust with your actual package path
)

// AddCoupon adds a new coupon to the database
func AddCoupon(c echo.Context) error {
	code := c.FormValue("code")
	description := c.FormValue("description")
	discount, _ := strconv.Atoi(c.FormValue("discount_percentage"))
	validFrom := c.FormValue("valid_from")
	expiresAt := c.FormValue("expires_at")

	ctx, queries := InitDB()
	_, err := queries.Exec(ctx, "INSERT INTO coupons (code, description, discount_percentage, valid_from, expires_at) VALUES ($1, $2, $3, $4, $5)", code, description, discount, validFrom, expiresAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add coupon"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Coupon added successfully"})
}

// ListCoupons retrieves all coupons from the database
func ListCoupons(c echo.Context) error {
	ctx, queries := InitDB()
	rows, err := queries.Query(ctx, "SELECT id, code, description, discount_percentage, valid_from, expires_at FROM coupons")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve coupons"})
	}
	defer rows.Close()
	var coupons []Coupon
	for rows.Next() {
		var c Coupon
		if err := rows.Scan(&c.ID, &c.Code, &c.Description, &c.DiscountPercentage, &c.ValidFrom, &c.ExpiresAt); err != nil {
			continue
		}
		coupons = append(coupons, c)
	}
	return c.JSON(http.StatusOK, coupons)
}
