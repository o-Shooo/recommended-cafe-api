package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"googlemaps.github.io/maps"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func init() {
	err := godotenv.Load()
	check(err)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/search", fetchCafeInfo)

	e.Logger.Fatal(e.Start(":3000"))
}

func fetchCafeInfo(c echo.Context) error {
	apiKey := os.Getenv("GOOGLE_API_KEY")

	mapsClient, err := maps.NewClient(maps.WithAPIKey(apiKey))
	check(err)

	//TODO:　radiusとLocationはフロント側から設定をもらう
	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{Lat: 35.7348, Lng: 139.7077},
		Radius:   1000,
		Keyword:  "cafe",
	}

	results, err := mapsClient.NearbySearch(c.Request().Context(), r)
	check(err)

	return c.JSON(http.StatusOK, results)

}
