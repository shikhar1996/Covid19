package database

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	State       string `json:"State" example:"Madhya Pradesh"`
	Count       int64  `json:"Count" example:"274405"`
	LastUpdated string `json:"LastUpdated" example:"2021-03-21T12:40:03.823+05:30"`
}

// Get count and last update time fof a given state in India
func GetCount(state string) (Response, error) {

	// Response vatiable
	var queryResponse Response

	// Handling of specific cases
	if (state == "Dadra and Nagar Haveli") || (state == "Daman and Diu") {
		state = "Dadra and Nagar Haveli and Daman and Diu"
	}
	conn, err := ConnectRedis()
	if err != nil {
		// zap.String("Error: Redis", err.Error())
	}
	count1, err := redis.Int64(conn.Do("HGET", state, "count"))
	if err == nil {
		queryResponse.State = state
		queryResponse.Count = count1
		lastUpdated1, err := redis.String(conn.Do("HGET", state, "lastUpdated"))
		queryResponse.LastUpdated = lastUpdated1
		return queryResponse, err
	}
	client, err := InitiateMongoClient()
	if err != nil {
		// zap.String("Error: Database Connection", err.Error())
		return queryResponse, err
	}

	// zap.String("Logger", state+" not present in cache")
	// Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(TABLE)

	// Resultant bson object
	var result bson.M

	if err = collection.FindOne(context.TODO(), bson.M{"_id": state}).Decode(&result); err != nil {
		// zap.String("Error: Data not Found", err.Error())
		return queryResponse, err
	}

	count := result["total"].(int64)
	lastUpdated := result["time"].(string)

	// Store the values in response struct
	queryResponse.State = state
	queryResponse.Count = count
	queryResponse.LastUpdated = lastUpdated

	// Add key(state) and values(count, lastupdated) to Redis
	_, err = conn.Do("HMSET", state, "count", count, "lastUpdated", lastUpdated)
	if err != nil {
		// zap.String("Error: Redis", err.Error())
	}

	// Add time to live (30 minutes = 1800 seconds)
	_, err = conn.Do("EXPIRE", state, 1800)
	if err != nil {
		// zap.String("Error: Redis", err.Error())
	}

	return queryResponse, nil
}
