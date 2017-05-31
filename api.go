package itunessearch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"log"
)

type SearchResult struct {
	ArtistId               float64 `json:"artistId"`
	ArtistName             string  `json:"artistName"`
	ArtistViewUrl          string  `json:"artistViewUrl"`
	ArtworkUrl100          string  `json:"artworkUrl100"`
	ArtworkUrl30           string  `json:"artworkUrl30"`
	ArtworkUrl60           string  `json:"artworkUrl60"`
	CollectionCensoredName string  `json:"collectionCensoredName"`
	CollectionExplicitness string  `json:"collectionExplicitness"`
	CollectionId           float64 `json:"collectionId"`
	CollectionName         string  `json:"collectionName"`
	CollectionPrice        float64 `json:"collectionPrice"`
	CollectionViewUrl      string  `json:"collectionViewUrl"`
	Country                string  `json:"country"`
	Currency               string  `json:"currency"`
	DiscCount              float64 `json:"discCount"`
	DiscNumber             float64 `json:"discNumber"`
	Kind                   string  `json:"kind"`
	PreviewUrl             string  `json:"previewUrl"`
	PrimaryGenreName       string  `json:"primaryGenreName"`
	RadioStationUrl        string  `json:"radioStationUrl"`
	ReleaseDate            string  `json:"releaseDate"`
	TrackCensoredName      string  `json:"trackCensoredName"`
	TrackCount             float64 `json:"trackCount"`
	TrackExplicitness      string  `json:"trackExplicitness"`
	TrackId                float64 `json:"trackId"`
	TrackName              string  `json:"trackName"`
	TrackNumber            float64 `json:"trackNumber"`
	TrackPrice             float64 `json:"trackPrice"`
	TrackTimeMillis        float64 `json:"trackTimeMillis"`
	TrackViewUrl           string  `json:"trackViewUrl"`
	WrapperType            string  `json:"wrapperType"`
}

type LookupResult struct {
	ArtistViewUrl												string		`json:"artistViewUrl"`
	ArtworkUrl60												string		`json:"artworkUrl60"`
	ArtworkUrl100												string		`json:"artworkUrl100"`
	IpadScreenshotUrls									[]string	`json:"ipadScreenshotUrls"`
	AppletvScreenshotUrls								[]string	`json:"appletvScreenshotUrls"`
	ArtworkUrl512												[]string	`json:"artworkUrl512"`
	IsGameCenterEnabled									bool			`json:"isGameCenterEnabled"`
	Features														[]string	`json:"features"`
	Kind																string		`json:"kind"`
	SupportedDevices										[]string	`json:"supportedDevices"`
	ScreenshotUrls											[]string	`json:"screenshotUrls"`
	Advisories													[]string	`json:"advisories"`
	TrackCensoredName										string		`json:"trackCensoredName"`
	TrackViewUrl												string		`json:"trackViewUrl"`
	ContentAdvisoryRating								string		`json:"contentAdvisoryRating"`
	LanguageCodesISO2A									[]string	`json:"languageCodesISO2A"`
	FileSizeBytes												string		`json:"fileSizeBytes"`
	SellerUrl														string		`json:"sellerUrl"`
	TrackContentRating									string		`json:"trackContentRating"`
	Currency														string		`json:"currency"`
	WrapperType													string		`json:"wrapperType"`
	Version															string		`json:"version"`
	ArtistId														int				`json:"artistId"`
	ArtistName													string		`json:"artistName"`
	Genres															[]string	`json:"genres"`
	Price																float64		`json:"price"`
	Description													string		`json:"description"`
	TrackId															int				`json:"trackId"`
	TrackName														string		`json:"trackName"`
	BundleId														string		`json:"bundleId"`
	ReleaseNotes												string		`json:"releaseNotes"`
	PrimaryGenreName										string		`json:"primaryGenreName"`
	IsVppDeviceBasedLicensingEnabled		bool			`json:"isVppDeviceBasedLicensingEnabled"`
	ReleaseDate													string		`json:"releaseDate"`
	FormattedPrice											string		`json:"formattedPrice"`
	MinimumOsVersion										string		`json:"minimumOsVersion"`
	PrimaryGenreId											int				`json:"primaryGenreId"`
	SellerName													string		`json:"sellerName"`
	GenreIds														[]string	`json:"genreIds"`
	CurrentVersionReleaseDate						string		`json:"currentVersionReleaseDate"`
}

func Search(query, country, media string) ([]SearchResult, error) {
	u := url.Values{}
	u["term"] = []string{query}
	u["country"] = []string{country}
	u["media"] = []string{media}
	res, err := http.Get("https://itunes.apple.com/search?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret struct {
		SearchResult []SearchResult
	}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret.SearchResult, nil
}

// New Lookup API
// search_term will be id, amgArtistId, upc, isbn
func Lookup(search_term string, search_term_value string, entity string, limit int, sort string) ([]LookupResult, error) {
	u := url.Values{}
	u[search_term] = []string{search_term_value}
	log.Println("URL: ", u.Encode())
	res, err := http.Get("https://itunes.apple.com/lookup?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret struct {
		LookupResult []LookupResult
	}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret.LookupResult, nil
}
