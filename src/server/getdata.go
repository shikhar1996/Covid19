package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shikhar1996/Covid19/src/database"
)

const api = "https://api.rootnet.in/covid19-in/stats/latest"

//update data in mongodb
func Updatedata(data []database.CovidDatabase) error {
	insertableList := make([]interface{}, len(data))
	for i, v := range data {
		insertableList[i] = v
	}
	//Get MongoDB connection using connectionhelper.
	client := database.InitiateMongoClient()

	//Create a handle to the respective collection in the database.
	collection := client.Database(database.DB).Collection(database.ISSUES)
	//Perform InsertMany operation & validate against the error.
	_, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil

}

//function to get covid data from public API
func Getdata() []database.CovidDatabase {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	data := result["data"].(map[string]interface{})["regional"].([]interface{})
	var covid []database.CovidDatabase
	s, _ := json.Marshal(data)
	json.Unmarshal([]byte(s), &covid)
	fmt.Printf("Objects : %+v", covid)
	/*
		for key, value := range covid {
			// Each value is an interface{} type, that is type asserted as a string
			fmt.Println(key, value)
		}
	*/
	//Convert the body to type string
	// sb := string(body)
	// jsonFile, err := os.Open(")
	// log.Printf(sb)
	return covid
}
