package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	constant "github.com/sahildhingraa/invidiousAPI/Constant"
	database "github.com/sahildhingraa/invidiousAPI/Database"
	playlist "github.com/sahildhingraa/invidiousAPI/Models"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	app := fiber.New()
	app.Get("/playlist/:PLID", fetchPlaylist)
	app.Get("/video/:VID", fetchVideo)

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(constant.PORT))
}

func fetchPlaylist(c *fiber.Ctx) error {

	id := c.Params("PLID")
	uri := constant.PlaylistURL + id + "?fields=" + constant.PlaylistFields

	response, err := http.Get(uri)
	Error(c, err, "Error fetching data from the external API")
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	Error(c, err, "Error reading response from the external API")

	var data playlist.Playlist
	err = json.Unmarshal(body, &data)
	Error(c, err, "Error parsing data to JSON")

	for i := range data.Videos {
		for _, thumbnail := range data.Videos[i].VideoThumbnail {
			if thumbnail.Quality == "maxres" {
				data.Videos[i].VideoThumbnail = []playlist.Thumbnail{thumbnail}
				break
			}
		}
	}
	return c.JSON(data)
}
func fetchVideo(c *fiber.Ctx) error {
	id := c.Params("VID")
	uri := constant.VideoURL + id + "?fields=" + constant.VideoFields

	response, err := http.Get(uri)
	Error(c, err, "Error fetching data from the external API")
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	Error(c, err, "Error reading response from the external API")

	var data playlist.Video
	err = json.Unmarshal(body, &data)
	Error(c, err, "Error parsing data to JSON")

	for i := range data.VideoThumbnail {
		if data.VideoThumbnail[i].Quality == "maxres" {
			data.VideoThumbnail = []playlist.Thumbnail{data.VideoThumbnail[i]}
			break
		}
	}
	for i := range data.FormatStreams {
		if data.FormatStreams[i].Itag == "22" {
			data.FormatStreams = []playlist.VideoFormat{data.FormatStreams[i]}
			break
		}
	}
	return c.JSON(data)
}
func Video(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := database.Mg.Db.Collection("Videos").Find(c.Context(), query)
	Error(c, err, err.Error())

	var videos []playlist.Video = make([]playlist.Video, 0)
	err = cursor.All(c.Context(), &videos)
	Error(c, err, err.Error())

	return c.JSON(videos)
}

func Error(c *fiber.Ctx, err error, message string) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(message)
	}
	return nil
}
