package itunessearch

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type SearchMap struct {
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

type SearchResult struct {
	ResultCount				int						`json:"resultCount"`
	Results						[]SearchMap		`json:"results"`
}

type LookupMap struct {
	Advisories													[]string	`json:"advisories"`
	AppletvScreenshotUrls								[]string	`json:"appletvScreenshotUrls"`
	ArtistId														float64		`json:"artistId"`
	ArtistName													string		`json:"artistName"`
	ArtistViewUrl												string		`json:"artistViewUrl"`
	ArtworkUrl60												string		`json:"artworkUrl60"`
	ArtworkUrl100												string		`json:"artworkUrl100"`
	ArtworkUrl512												string	`json:"artworkUrl512"`
	BundleId														string		`json:"bundleId"`
	ContentAdvisoryRating								string		`json:"contentAdvisoryRating"`
	Currency														string		`json:"currency"`
	CurrentVersionReleaseDate						string		`json:"currentVersionReleaseDate"`
	Description													string		`json:"description"`
	Features														[]string	`json:"features"`
	FileSizeBytes												string		`json:"fileSizeBytes"`
	FormattedPrice											string		`json:"formattedPrice"`
	Genres															[]string	`json:"genres"`
	GenreIds														[]string	`json:"genreIds"`
	IpadScreenshotUrls									[]string	`json:"ipadScreenshotUrls"`
	IsGameCenterEnabled									bool			`json:"isGameCenterEnabled"`
	IsVppDeviceBasedLicensingEnabled		bool			`json:"isVppDeviceBasedLicensingEnabled"`
	Kind																string		`json:"kind"`
	LanguageCodesISO2A									[]string	`json:"languageCodesISO2A"`
	MinimumOsVersion										string		`json:"minimumOsVersion"`
	Price																float64		`json:"price"`
	PrimaryGenreId											float64		`json:"primaryGenreId"`
	PrimaryGenreName										string		`json:"primaryGenreName"`
	ReleaseDate													string		`json:"releaseDate"`
	ReleaseNotes												string		`json:"releaseNotes"`
	ScreenshotUrls											[]string	`json:"screenshotUrls"`
	SellerName													string		`json:"sellerName"`
	SellerUrl														string		`json:"sellerUrl"`
	SupportedDevices										[]string	`json:"supportedDevices"`
	TrackCensoredName										string		`json:"trackCensoredName"`
	TrackContentRating									string		`json:"trackContentRating"`
	TrackId															float64		`json:"trackId"`
	TrackName														string		`json:"trackName"`
	TrackViewUrl												string		`json:"trackViewUrl"`
	Version															string		`json:"version"`
	WrapperType													string		`json:"wrapperType"`
}

type LookupResult struct {
	ResultCount				int						`json:"resultCount"`
	Results						[]LookupMap		`json:"results"`
}

func Search(query, country, media string) (*SearchResult, error) {
	u := url.Values{}
	u["term"] = []string{query}
	u["country"] = []string{country}
	u["media"] = []string{media}
	res, err := http.Get("https://itunes.apple.com/search?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret := SearchResult{}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// New Lookup API
// search_term will be id, amgArtistId, upc, isbn
func Lookup(search_term string, search_term_value string, entity string, limit string, sort string) (*LookupResult, error) {
	u := url.Values{}
	u[search_term] = []string{search_term_value}
	u["entity"] = []string{entity}
	u["limit"] = []string{limit}
	u["sort"] = []string{sort}
	res, err := http.Get("https://itunes.apple.com/lookup?" + u.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret := LookupResult{}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
