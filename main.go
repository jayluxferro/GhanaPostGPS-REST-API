package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	gp "github.com/jayluxferro/ghanapostgps"
)

var params gp.Params

const identifier = "CenterLatitude"
const getAddressIdentifier = "PostCode"

type DataResponse struct {
	Table []Info
}

type AddressResponse struct {
	Table []Address
}

type Address struct {
	GPSName  string
	Region   string
	District string
	PostCode string
	NLat     float64
	Slat     float64
	WLong    float64
	Elong    float64
	Area     string
	Street   string
}

type Info struct {
	Area            string
	CenterLatitude  float64
	CenterLongitude float64
	District        string
	EastLat         float64
	EastLong        float64
	GPSName         string
	NorthLat        float64
	NorthLong       float64
	PostCode        string
	Region          string
	SouthLat        float64
	SouthLong       float64
	Street          string
	WestLat         float64
	WestLong        float64
}

func unAuthorized(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
}

func responseData(c *gin.Context, found bool, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"found": found,
	})
}

func getAddressHandler(c *gin.Context) {
	var dataResponse AddressResponse
	lat := string(c.PostForm("lat"))
	long := string(c.PostForm("long"))

	if !(len(lat) > 0 && len(long) > 0) {
		unAuthorized(c)
		return
	}

	response := gp.GetAddress(lat, long, &params)
	if !strings.Contains(response, getAddressIdentifier) {
		responseData(c, false, dataResponse)
		return
	}

	response = convertToJSON(response)

	json.Unmarshal([]byte(response), &dataResponse)
	responseData(c, true, dataResponse)
}

func getLocationHandler(c *gin.Context) {
	var dataResponse DataResponse
	isValid, address := gp.IsValidGPAddress(c.PostForm("address"))
	if !isValid {
		unAuthorized(c)
		return
	}

	response := gp.GetLocation(address, &params)

	if !strings.Contains(response, identifier) {
		responseData(c, false, dataResponse)
		return
	}

	response = convertToJSON(response)

	json.Unmarshal([]byte(response), &dataResponse)
	responseData(c, true, dataResponse)
}

func convertToJSON(data string) string {
	in := []byte(data)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	out, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

func main() {
	// inits
	prefix := "GPGPS_"
	params = gp.Params{}
	params.AndroidCert = os.Getenv(prefix + "androidCert")
	params.AndroidPackage = os.Getenv(prefix + "androidPackage")
	params.ApiKey = os.Getenv(prefix + "apiKey")
	params.ApiURL = os.Getenv(prefix + "apiURL")
	params.AsaaseAPI = os.Getenv(prefix + "asaaseAPI")
	params.Country = os.Getenv(prefix + "country")
	params.CountryName = os.Getenv(prefix + "countryName")
	params.Language = os.Getenv(prefix + "language")
	params.LanguageCode = os.Getenv(prefix + "languageCode")
	params.UUID = os.Getenv(prefix + "uuid")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://sperixlabs.org")
	})

	// routes
	router.POST("/", getLocationHandler)
	router.POST("/get-location", getLocationHandler)
	router.POST("/get-address", getAddressHandler)

	router.Run(":" + port)
}
