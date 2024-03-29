package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDB() error {
	// uri := os.Getenv("MONGODB_URI")
	uri:= viper.GetString("mongo.uri")
	if uri == "" {
		return errors.New("you must set your mongo uri environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully connected to Mongodb Instance")
	}

	db = client.Database(viper.GetString("mongo.database"))

	return nil
}

// func CloseDB() error {
// 	return db.Client().Disconnect(context.Background())
// }
