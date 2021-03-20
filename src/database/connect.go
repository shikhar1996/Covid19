package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CovidDatabase struct {
	State      string `json:"loc"`
	Indian     int    `json:"confirmedCasesIndian"`
	Foreigner  int    `json:"confirmedCasesForeign"`
	Discharged int    `json:"discharged"`
	Deaths     int    `json:"deaths"`
	Total      int    `json:"totalConfirmed"`
	// Time time.Time `json:"time"`
}

const (
	CONNECTIONSTRING = "mongodb+srv://shikhar:shikhar%4012@cluster0.mya9h.mongodb.net/covidDatabase?retryWrites=true&w=majority"
	DB               = "covidDatabase"
	ISSUES           = "India"
)

func InitiateMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client
	uri := CONNECTIONSTRING
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}
