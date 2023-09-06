package tracker

type Er struct {
	Status int
	Text   string
}

type Artist struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstalbum"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

type Relation struct {
	Index []struct {
		Id             uint64              `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}

var (
	Artists   []Artist
	Relations Relation
)

type Alldata struct {
	AllArtist   []Artist
	FoundArtist []Artist
}

var (
	ArtistApi  string = "https://groupietrackers.herokuapp.com/api/artists"
	ArtistDate string = "https://groupietrackers.herokuapp.com/api/relation"
)
