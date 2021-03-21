package geoencoding

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type Coordinates struct {
	Lat  string
	Long string
}

const (
	URL    = "http://api.positionstack.com"
	APIKEY = "9dc26181b71942060fcf15948628ce3a"
)

// Get Indian state from the GPS coordinate
func GetState(coordinate Coordinates) (string, error) {
	lat1 := coordinate.Lat
	long1 := coordinate.Long
	coordinateParam := lat1 + "," + long1

	baseURL, _ := url.Parse(URL)

	baseURL.Path += "v1/reverse"

	params := url.Values{}
	// Access Key
	params.Add("access_key", APIKEY)
	// Query = latitude,longitude
	params.Add("query", coordinateParam)
	// Optional parameters
	params.Add("output", "json")

	baseURL.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", baseURL.String(), nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		zap.String("Error Connection", err.Error())
		return "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	// Get data field from response
	data := result["data"].([]interface{})[0].(map[string]interface{})
	country := data["country"]
	region := data["region"]

	// Check for country
	if country != "India" {
		// fmt.Println("Country is not India")
		return "Country is not India", errors.New("Coordinates are not of India")
	}

	return region.(string), nil
}
