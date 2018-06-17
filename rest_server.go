package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
)

type User struct {
	Id  string `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

var m map[string]User

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	u, found := m[id]
	if !found {
		fmt.Println("Not found")
		return c.JSON(http.StatusNoContent, "not found")
	}
	return c.JSON(http.StatusOK, u)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	m[u.Id] = *u
	return c.JSON(http.StatusCreated, m[u.Id])
	// or
	// return c.XML(http.StatusCreated, u)
}

func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	delete(m, id)
	return c.JSON(http.StatusOK, m[id])
}

func main() {
	m = make(map[string]User)
	e := echo.New()
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
