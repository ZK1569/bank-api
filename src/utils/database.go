package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Coll *mongo.Collection
}

func ConnectionToBd() *Database {
	uri := EnvVariable.DatabaseUri
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("Mimir").Collection("User")

	return &Database{
		Coll: coll,
	}
}

// defer func() {
// 	if err := client.Disconnect(ctx); err != nil {
// 		log.Fatal(err)
// 	}
// }()
