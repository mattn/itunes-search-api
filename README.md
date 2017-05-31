# iTunes Search API
iTunes Search API for Go.

# Installation
```golang
go get github.com/Fyb3roptik/itunes-search-api
```

# Examples
```golang
import(
  "github.com/Fyb3roptik/itunes-search-api"
)
```
# Search Example
```golang
results, err := itunessearch.Search("shinedown", "US", "music")
if err != nil {
	log.Fatal(err)
}
```
# Lookup Example
```golang
lookup, err := itunessearch.Lookup("id", "1202948413", "", "1", "")
if err != nil {
	log.Fatal(err)
}
```
