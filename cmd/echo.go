package main

import (
	"hello/fizzbuzz"
	"hello/oscar"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/oscarmale", oscarmale)
	e.GET("/fizzbuzz/:number", fizzbuzzHandler)
	e.POST("/fizzbuzz", postFizzbuzzHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func oscarmale(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func fizzbuzzHandler(c echo.Context) error {
	numberString := c.Param("number")
	n, _ := strconv.Atoi(numberString)
	return c.String(http.StatusOK, fizzbuzz.Say(n))
}

func postFizzbuzzHandler(c echo.Context) error {

	var req map[string]int
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	type FizzbuzzResponse struct {
		Number   int    `json:"number"`
		FizzBuzz string `json:"fizzbuzz"`
	}

	return c.JSON(http.StatusOK, FizzbuzzResponse{
		Number:   req["number"],
		FizzBuzz: fizzbuzz.Say(req["number"]),
	})
}
