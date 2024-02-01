package main

import (
	"encoding/xml"
	"net/http"

	"log"

	"github.com/labstack/echo/v4"
)

type Team struct {
	XMLName xml.Name `xml:"team"`
	Name    string   `xml:"name,attr"`
	Member  Member   `xml:"member"`
}

type Member struct {
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name,attr"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/show", show)
	e.GET("/users/:id", getUser)
	e.POST("/save", save)

	err := e.Start(":1323")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from the path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.Get("/show", show)
func show(c echo.Context) error {
	//Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	//format response as xml with seperate root element

	t := &Team{
		Name: team,
		Member: Member{
			Name: member,
		},
	}

	return c.XML(http.StatusOK, t)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
