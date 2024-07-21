package handler

import (
	"strconv"

	"example.com/test/view/components"
	"github.com/labstack/echo/v4"
)

func GetBookingView(c echo.Context) error {
	ctx, queries := InitDB()
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return RenderTemplComponent(c, components.ErrorPage("Invalid movie ID"))
	}

	movie, err := queries.GetMovie(ctx, int32(movieID))
	if err != nil {
		return RenderTemplComponent(c, components.ErrorPage("Movie not found"))
	}
	title := movie.Title.String
	return RenderTemplComponent(c, components.SeatEditor(12, 24, c.Param("id"), title))
}

func GetMovieInfo(c echo.Context) error {
	ctx, queries := InitDB()
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return RenderTemplComponent(c, components.ErrorPage("Invalid movie ID"))
	}

	movie, err := queries.GetMovie(ctx, int32(movieID))
	if err != nil {
		return RenderTemplComponent(c, components.ErrorPage("Movie not found"))
	}

	return RenderTemplComponent(c, components.MovieInfo(movie))
}

func DeleteMovie(c echo.Context) error {
	// Extract movie_id from the route parameter
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// If the movie_id is not a valid integer, return an error response
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid movie ID"})
	}

	// Initialize database context and prepared queries
	ctx, queries := InitDB()

	// Execute the delete operation
	if err := queries.DeleteMovie(ctx, int32(movieID)); err != nil {
		// If there is an error during the delete operation, return an internal server error response
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete movie"})
	}

	// If the delete operation is successful, return a success response
	return c.JSON(http.StatusOK, echo.Map{"message": "Movie deleted successfully"})
}