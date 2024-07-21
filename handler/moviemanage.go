package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"example.com/test/db"
)

// CreateMovie xử lý yêu cầu POST để tạo phim mới
func CreateMovie(c echo.Context) error {
	var movie db.Movie
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx, queries := InitDB()
	newMovie, err := queries.CreateMovie(ctx, movie.Title, movie.Overview, movie.OriginalLanguage, movie.AgeRes, movie.ReleaseDate, movie.Status, movie.Tagline, movie.Length, movie.Url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMovie)
}

// UpdateMovie xử lý yêu cầu PUT để cập nhật phim hiện có
func UpdateMovie(c echo.Context) error {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid movie ID")
	}
	var movie db.Movie
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx, queries := InitDB()
	err = queries.UpdateMovie(ctx, movie.Title, movie.Overview, movie.OriginalLanguage, movie.AgeRes, movie.ReleaseDate, movie.Status, movie.Tagline, movie.Length, movie.Url, int32(movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Phim đã được cập nhật thành công")
}

// DeleteMovie xử lý yêu cầu DELETE để xóa phim
func DeleteMovie(c echo.Context) error {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid movie ID")
	}
	ctx, queries := InitDB()
	err = queries.DeleteMovie(ctx, int32(movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Phim đã được xóa thành công")
}
