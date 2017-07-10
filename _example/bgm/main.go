package main

import (
	"fmt"
	"log"

	"github.com/mattn/itunes-search-api"
)

func main() {
	results, err := itunessearch.Search("マツケンサンバ", "JP", "music")
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results.Results {
		fmt.Println(result.ArtistName, result.TrackName, result.CollectionViewUrl)
		playURL(result.PreviewUrl)
	}
}
