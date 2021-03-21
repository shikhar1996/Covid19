package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/shikhar1996/Covid19/src/database"
	"github.com/shikhar1996/Covid19/src/geoencoding"
	"go.uber.org/zap"

	_ "github.com/shikhar1996/Covid19/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var Message string

// echo http://localhost:1323/Cases?latitude=21.903132&longitude=77.912053
// default value Betul Coordinates :)
// GetDataByState godoc
// @Summary Get total number of covid cases
// @Description Take the latitude and longitude as input and return total number of coved cases in that state along with time stamp
// @Tags covid
// @Accept  json
// @Produce  json
// @Param latitude query string true "Latitude"
// @Param longitude query string true "Longitude"
// @Success 200 {object} database.Response
// @Failure 400 {string} Message
// @Router /total_count [post]
func getCovidCount(c echo.Context) error {

	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")
	var coordinates geoencoding.Coordinates
	coordinates.Lat = latitude
	coordinates.Long = longitude

	// zap.String("Input Coordinates", "("+latitude+", "+longitude+")")
	if latitude == "" || longitude == "" {
		Message = "Invalid Coordinates"
		return echo.NewHTTPError(http.StatusBadRequest, Message)
	}
	state, err := geoencoding.GetState(coordinates)

	if err != nil {
		zap.String("Error : Invalid Input", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response, err := database.GetCount(state)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

// e.GET("/Data", updateDatabase)
// Update Database godoc
// @Summary Update the MongoDB Database
// @Description This API calls "https://api.rootnet.in/covid19-in/stats/latest" API and update the Covid Data
// @Tags covid
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /update_data [get]
func updateDatabase(c echo.Context) error {

	data := database.Getdata()
	err := database.Updatedata(data)
	if err != nil {
		zap.String("Error: Database Updation", err.Error())
		c.String(http.StatusInternalServerError, "Database not Updated")
	}

	return c.String(http.StatusOK, "Database Updated")

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// @title Swagger API for Covid India Data
// @version 1.0
// @description This is a server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email shikhar.agrawal789@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes https
func Redirect() {

	// Echo instance
	e := echo.New()

	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	// e.Use(mw.Logger())
	// Check if server is running or not
	e.GET("/", HealthCheck)

	// Swagger API
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Normal API
	e.POST("/total_count", getCovidCount)
	e.GET("/update_data", updateDatabase)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
	fmt.Println("Port is:", e.Listener.Addr().(*net.TCPAddr).Port)

}
