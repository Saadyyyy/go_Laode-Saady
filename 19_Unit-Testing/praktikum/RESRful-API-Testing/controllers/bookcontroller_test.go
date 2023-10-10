package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooksController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidTableName", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/wrongtable", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.Error(t, GetAllBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestCreateBookController(t *testing.T) {
	e := echo.New()
	t.Run("InputValid", func(t *testing.T) {
		reqBody := `{
			"title": "Petualangan Sang Pemberani",
			"author": "Rizky Pratama",
			"publisher": "Gema Pustaka"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InputValid")

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestGetBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "Cinta dan Kenangan",
			"author": "Amanda Putri",
			"publisher": "Sentosa Media"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/2", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "Cinta dan Kenangan",
			"author": "Amanda Putri",
			"publisher": "Sentosa Media"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/2", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/3", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("3")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/4", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("4")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
