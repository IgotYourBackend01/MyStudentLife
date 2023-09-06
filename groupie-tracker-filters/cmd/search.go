package tracker

import (
	"strconv"
	"strings"
)

var FoundArtist []Artist

func SearchArtist(info string) {
	info = strings.ToLower(info)
	var NewSearch []Artist

	for i, k := range Artists {
		for city, _ := range k.DatesLocations {
			if strings.Contains(strings.ToLower(city), info) {
				if IsCorrect(k.Id, NewSearch) {
					NewSearch = append(NewSearch, Artists[i])
				}
			}
		}
		if strings.Contains(strings.ToLower(k.Name), info) || strconv.Itoa(k.CreationDate) == info || strings.Contains(k.FirstAlbum, info) {
			if IsCorrect(k.Id, NewSearch) {
				NewSearch = append(NewSearch, Artists[i])
			}
		}
		for _, m := range k.Members {
			if strings.Contains(strings.ToLower(m), info) {
				if IsCorrect(k.Id, NewSearch) {
					NewSearch = append(NewSearch, Artists[i])
				}
			}
		}
	}
	FoundArtist = NewSearch
}

func IsCorrect(id int, info []Artist) bool {
	for _, m := range info {
		if id == m.Id {
			return false
		}
	}
	return true
}
