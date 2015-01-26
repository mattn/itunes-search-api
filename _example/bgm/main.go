package main

import (
	"fmt"
	"github.com/mattn/itunes-search-api"
	"log"
)

func main() {
	results, err := itunessearch.Search("マツケンサンバ", "JP", "music")
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result.ArtistName, result.TrackName, result.CollectionViewUrl)
		playURL(result.PreviewUrl)
	}
}
