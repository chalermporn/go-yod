package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"hello/fizzbuzz"
	"hello/oscar"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// Server header
	e.Use(ServerHeader)

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Routes
	e.GET("/oscarmale", oscarmale)
	e.GET("/fizzbuzz/:number", fizzbuzzHandler)
	e.POST("/fizzbuzz", postFizzBuzzHandler)
	e.POST("/token", tokenHandler)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	h := &randomFizzBuzz{random: r1}
	e.GET("/fizzbuzzr", h.handler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// type In

// Handler
func oscarmale(c echo.Context) error {
	result := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")
	return c.JSON(http.StatusOK, result)
}

func fizzbuzzHandler(c echo.Context) error {
	// numberString := c.Param("number")
	// n, _ := strconv.Atoi(numberString)
	// return c.String(http.StatusOK, fizzbuzz.Say(n))

	// tokenString := c.Request().Header.Get("Authorization")[7:]
	// type ErrorResponse struct {
	// 	Message string `json:"message"`
	// }

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		c.JSON(http.StatusInternalServerError, ErrorResponse{
	// 			Message: fmt.Sprintf("%s", r),
	// 		})
	// 	}
	// }()

	// var validateSignature = func(tokenString *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// }
	// _, err := jwt.Parse(tokenString, validateSignature)

	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, ErrorResponse{
	// 		Message: err.Error(),
	// 	})
	// }

	numberString := c.Param("number")
	n, _ := strconv.Atoi(numberString)
	// return c.String(http.StatusOK, fizzbuzz.New(n).String())
	type fizzbuzzResponse struct {
		Number  string `json:"number"`
		Message string `json:"message"`
	}

	return c.JSON(http.StatusOK, fizzbuzzResponse{
		Number:  numberString,
		Message: fizzbuzz.New(n).String(),
	})

}

type randomer interface {
	Intn(int) int
}

type randomFizzBuzz struct {
	random randomer
}

func (r *randomFizzBuzz) handler(c echo.Context) error {
	n := r.random.Intn(100)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"number":  n,
		"message": fizzbuzz.Say(n),
	})
}

func randomFizzBuzzHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, fizzbuzzController(rand.Intn(100)))
}

func fizzbuzzController(n int) map[string]interface{} {
	return map[string]interface{}{
		"number":  n,
		"message": fizzbuzz.Say(n),
	}
}

func postFizzBuzzHandler(c echo.Context) error {
	var req map[string]int
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	type fizzbuzzResponse struct {
		Number   int    `json:"number"`
		FizzBuzz string `json:"fizzbuzz"`
	}

	return c.JSON(http.StatusOK, fizzbuzzResponse{
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
