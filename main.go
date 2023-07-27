package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	constant "github.com/sahildhingraa/invidiousAPI/Constant"
	database "github.com/sahildhingraa/invidiousAPI/Database"
	model "github.com/sahildhingraa/invidiousAPI/Models"
)

func main() {
	app := fiber.New()
	app.Get("/playlist/:PLID", fetchPlaylist)
	app.Get("/video/:VID", fetchVideo)
	app.Get("/videos", getAllVideos)
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

	var data model.Playlist
	err = json.Unmarshal(body, &data)
	Error(c, err, "Error parsing data to JSON")

	for i := range data.Videos {
		for _, thumbnail := range data.Videos[i].VideoThumbnail {
			if thumbnail.Quality == "maxres" {
				data.Videos[i].VideoThumbnail = []model.Thumbnail{thumbnail}
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

	var data model.Video
	err = json.Unmarshal(body, &data)
	Error(c, err, "Error parsing data to JSON")

	for i := range data.VideoThumbnail {
		if data.VideoThumbnail[i].Quality == "maxres" {
			data.VideoThumbnail = []model.Thumbnail{data.VideoThumbnail[i]}
			break
		}
	}
	for i := range data.FormatStreams {
		if data.FormatStreams[i].Itag == "22" {
			data.FormatStreams = []model.VideoFormat{data.FormatStreams[i]}
			break
		}
	}
	database.InsertVideo(data)
	return c.JSON(data)
}
func getAllVideos(c *fiber.Ctx) error {
	result := database.GetAllVideos()
	return c.JSON(result)
}

func Error(c *fiber.Ctx, err error, message string) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(message)
	}
	return nil
}
