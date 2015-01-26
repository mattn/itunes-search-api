package itunessearch

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Result struct {
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

func Search(query, country, media string) ([]Result, error) {
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
		Results []Result
	}
	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret.Results, nil
}
