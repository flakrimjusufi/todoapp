package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoapp/controllers"
)

var createToDoList = `{"labels":"Don't forget the breakfast","comments":"You should take a coffee too","due_date":"2021-11-12T10:45:00Z"}
`

var updateToDoList = `{"ID":1,"labels":"Don't forget to eat something","comments":"You should take a coffee too","due_date":"2021-11-12T10:55:00Z"}
`

var toDoListNotFound = `"ToDoList not found!"
`

func TestCreateToDoList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createToDoList))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList")

	// Assertions
	if assert.NoError(t, controllers.CreateToDoList(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetToDoListById(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.GetToDoById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, true, rec.Body.Len() > 0)
	}
}

func TestUpdateToDoList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(updateToDoList))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.UpdateToDoById(c)) {
		if rec.Code == http.StatusNotFound {
			assert.Equal(t, toDoListNotFound, rec.Body.String())
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, updateToDoList, rec.Body.String())
		}
	}
}

func TestDeleteToDoList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(updateToDoList))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList/:id")
	c.SetParamNames("id")
	c.SetParamValues("10")

	// Assertions
	if assert.NoError(t, controllers.DeleteToDo(c)) {
		if rec.Code == http.StatusNotFound {
			assert.Equal(t, toDoListNotFound, rec.Body.String())
		} else {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	}
}

func TestGetToDoList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList")

	// Assertions
	if assert.NoError(t, controllers.GetAllToDos(c)) {
		if rec.Body.Len() == 0 {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, true, rec.Body.Len() > 0)
		}
	}
}
