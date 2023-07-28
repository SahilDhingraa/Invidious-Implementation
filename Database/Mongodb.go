package database

import (
	"context"
	"log"
	"os"

	constant "github.com/sahildhingraa/invidiousAPI/Constant"
	model "github.com/sahildhingraa/invidiousAPI/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ColName = "video"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	Error(err)
	collection = client.Database(constant.DbName).Collection(ColName)
}
func InsertVideo(video model.Video) interface{} {
	inserted, err := collection.InsertOne(context.Background(), video)
	Error(err)
	return inserted.InsertedID
}
func GetAllVideos() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	Error(err)
	defer cursor.Close(context.Background())

	var Videos []primitive.M
	for cursor.Next(context.Background()) {
		var video bson.M
		err := cursor.Decode(&video)
		Error(err)
		Videos = append(Videos, video)
	}
	return Videos
}

func Error(er error) {
	if er != nil {
		log.Fatal(er)
	}
}
