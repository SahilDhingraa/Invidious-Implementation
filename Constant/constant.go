package constant

// MAIN URI
const (
	BaseURL = "https://inv.in.projectsegfau.lt/api/v1/"
	PORT    = ":3000"
)

// PLAYLIST
const (
	PlaylistURL    = BaseURL + "playlists/"
	PlaylistFields = "playlistId,title,videos,videoCount"
)

// VIDEO
const (
	VideoURL    = BaseURL + "videos/"
	VideoFields = "videoId,title,videoThumbnails,lengthSeconds,formatStreams"
)

// THUMBNAIL
const (
	ThumbnailFields = "url,width,height"
)

// DATABASE
const (
	DbName   = "invidious-implementation"
	MongoURI = "mongodb://localhost:27017" + DbName
)
