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

// Enter user, password, database and collection details
const (
	CONNECTIONSTRING = "mongodb+srv://<user>:<password>@cluster0.mya9h.mongodb.net/<database>?retryWrites=true&w=majority"
	DB               = "<database>"
	TABLE            = "<collection>"
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
