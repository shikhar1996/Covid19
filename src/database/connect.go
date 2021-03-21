package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type CovidDatabase struct {
	State      string `json:"loc" bson:"_id,omitempty"`
	Indian     int64  `json:"confirmedCasesIndian"`
	Foreigner  int64  `json:"confirmedCasesForeign"`
	Discharged int64  `json:"discharged"`
	Deaths     int64  `json:"deaths"`
	Total      int64  `json:"totalConfirmed"`
	Time       string `json:"time"`
}

// Add you password and database name
const (
	CONNECTIONSTRING = "mongodb+srv://shikhar:<password>@cluster0.mya9h.mongodb.net/covidDatabase?retryWrites=true&w=majority"
	DB               = "covidDatabase"
	TABLE            = "India"
)

// Create connection
func InitiateMongoClient() (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	uri := CONNECTIONSTRING
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(0)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		zap.String("Error: Database Connection", err.Error())
	}
	return client, err
}
