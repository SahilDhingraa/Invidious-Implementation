package database

import (
	"context"
	"log"
	"time"

	constant "github.com/sahildhingraa/invidiousAPI/Constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

type Video struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	Title          string `json:"title"`
	VideoID        string `json:"videoID"`
	VideoThumbnail string `json:"videoThumbnail"`
	PlaylistID     string `json:"playlistID"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(constant.MongoURI))
	Error(err)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	Error(err)
	db := client.Database(constant.DbName)

	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func Error(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
