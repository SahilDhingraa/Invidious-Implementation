package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	Title      string  `json:"title"`
	PlaylistID string  `json:"playlistId"`
	Videos     []Video `json:"videos"`
	VideoCount int64   `json:"videoCount"`
}
type Video struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title          string        `json:"title"`
	VideoID        string        `json:"videoId"`
	VideoThumbnail []Thumbnail   `json:"videoThumbnails"`
	LengthSeconds  int64         `json:"lengthSeconds"`
	FormatStreams  []VideoFormat `json:"formatStreams"`
}
type Thumbnail struct {
	Quality string `json:"quality"`
	URL     string `json:"url"`
}
type VideoFormat struct {
	Url  string `json:"url"`
	Itag string `json:"itag"`
	Type string `json:"type"`
}
