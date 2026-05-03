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
	"github.com/joho/godotenv"
)

var params gp.Params

const identifier = "CenterLatitude"
const getAddressIdentifier = "PostCode"

type GetLocationRequest struct {
	Address string `json:"address" form:"address"`
}

type GetAddressRequest struct {
	Lat  string `json:"lat" form:"lat"`
	Long string `json:"long" form:"long"`
}

type DataResponse struct {
	Table []Info
}

type AddressResponse struct {
	Table []Address
}

type Address struct {
	GPSName   string
	Region    string
	District  string
	PostCode  string
	NLat      float64
	Slat      float64
	WLong     float64
	Elong     float64
	Area      string
	AddressV1 string
	Street    string
}

type Info struct {
	Area            string
	AddressV1       string
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

func bindInput(c *gin.Context, v interface{}) bool {
	ct := c.GetHeader("Content-Type")
	if strings.HasPrefix(ct, "application/json") {
		if err := c.ShouldBindJSON(v); err != nil {
			return false
		}
	} else {
		if err := c.ShouldBind(v); err != nil {
			return false
		}
	}
	return true
}

func getAPIKeysHandler(c *gin.Context) {
	defaults := params
	defaults.ApiURL = gp.BaseAPIURL
	c.JSON(http.StatusOK, gin.H{"data": gp.GetAPIKeys(&defaults)})
}

func getAddressHandler(c *gin.Context) {
	var dataResponse AddressResponse
	var req GetAddressRequest

	if !bindInput(c, &req) {
		unAuthorized(c)
		return
	}

	if !(len(req.Lat) > 0 && len(req.Long) > 0) {
		unAuthorized(c)
		return
	}

	response := gp.GetAddress(req.Lat, req.Long, &params)
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
	var req GetLocationRequest

	if !bindInput(c, &req) {
		unAuthorized(c)
		return
	}

	isValid, address := gp.IsValidGPAddress(req.Address)
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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}

func main() {
	_ = godotenv.Load(".env")

	prefix := "GPGPS_"
	params = gp.Params{}
	params.ApiURL = os.Getenv(prefix + "apiURL")
	params.Authorization = os.Getenv(prefix + "authorization")
	params.AsaaseUser = os.Getenv(prefix + "asaaseUser")
	params.LanguageCode = os.Getenv(prefix + "languageCode")
	params.Language = os.Getenv(prefix + "language")
	params.DeviceId = os.Getenv(prefix + "deviceId")
	params.AndroidCert = os.Getenv(prefix + "androidCert")
	params.AndroidPackage = os.Getenv(prefix + "androidPackage")
	params.CountryName = os.Getenv(prefix + "countryName")
	params.Country = os.Getenv(prefix + "country")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	var production = false
	if os.Getenv("MODE") == "prod" {
		production = true
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"production": production,
		})
	})

	router.POST("/", getLocationHandler)
	router.POST("/get-location", getLocationHandler)
	router.POST("/get-address", getAddressHandler)
	router.POST("/api-keys", getAPIKeysHandler)

	router.Run(":" + port)
}
