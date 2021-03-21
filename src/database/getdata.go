package database

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const API = "https://api.rootnet.in/covid19-in/stats/latest"

// Update data in mongodb
func Updatedata(data []CovidDatabase) error {

	insertableList := make([]interface{}, len(data))
	curTime := time.Now().String()

	for i, v := range data {
		v.Time = curTime
		insertableList[i] = v
	}

	// Get MongoDB connection using connectionhelper.
	client, err := InitiateMongoClient()
	if err != nil {
		// zap.String("Error: Database Connection", err.Error())
	}

	// Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TABLE)

	// Perform InsertMany operation & validate against the error.
	// Drop collection if exists
	// It might be possible that server is updating and an user is requesting the data at the
	// time. If we want to make it failsafe we can add replica of the same collection which will
	// be updated at different time.
	if err = collection.Drop(context.TODO()); err != nil {
		// zap.String("Error: Table Drop", err.Error())
	}
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		// zap.String("Error: MongoDB", err.Error())
		return err
	}
	// Return success without any error.
	return nil
}

// Function to get covid data from public API
func Getdata() []CovidDatabase {
	resp, err := http.Get(API)
	if err != nil {
		// zap.String("Error: Getting API data", err.Error())
	}
	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// zap.String("Error: Reading API data", err.Error())
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	// Store all relevant objects in data
	data := result["data"].(map[string]interface{})["regional"].([]interface{})

	var covid []CovidDatabase
	s, _ := json.Marshal(data)
	json.Unmarshal([]byte(s), &covid)

	// fmt.Printf("Objects : %+v", covid)

	return covid
}
