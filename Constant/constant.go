package constant

const (
	BaseURL = "https://inv.in.projectsegfau.lt/api/v1/"
)

const (
	PlaylistURL    = BaseURL + "playlists/"
	PlaylistFields = "playlistId,title,videos,videoCount"
)
const (
	VideoURL    = BaseURL + "videos/"
	VideoFields = "videoId,title,videoThumbnails,lengthSeconds,formatStreams"
)
const (
	ThumbnailFields = "url,width,height"
)
