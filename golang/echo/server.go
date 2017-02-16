package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"name"`
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	//	e.POST("/users", saveUser)
	//	e.GET("/users/:id", getUser)
	//	e.PUT("/users/:id", updateUser)
	//	e.DELETE("/users/:id", deleteUser)

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.POST("/users/new", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})

	e.GET("/json", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
