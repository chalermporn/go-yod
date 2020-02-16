package main

import (
	"fmt"
	"hello/fizzbuzz"
	"hello/oscar"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	e.POST("/token", tokenHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func oscarmale(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func fizzbuzzHandler(c echo.Context) error {

	// tokenString := c.Request().Header.Get("Authorization")[7:]
	type ErrorResponse struct {
		Message string `json:"message"`
	}

	defer func(){
		if r := recover(); r != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			})
		}
	}


	var validateSignature = func(tokenString *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	}
	_, err := jwt.Parse(tokenString, validateSignature)
	
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: err.Error(),
		})
	}

	fmt.Print("tokenString :", tokenString)

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

func tokenHandler(c echo.Context) error {

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		Issuer:    "bird",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	// fmt.Printf("%v %v", ss, err)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	// return c.String(http.StatusOK, ss)
	type TokenResponse struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: ss,
	})

}
