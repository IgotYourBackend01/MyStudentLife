package tracker

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var FiltredArtists []Artist

func FilterData(minCreate, maxCreate, minFirst, maxFirst, location string, membersNum []string) {
	var NewData []Artist
	fmt.Println(membersNum)
	for i, artist := range Artists {
		if IsUnique(artist.Id, NewData) && CreationDateCheck(minCreate, maxCreate, artist) && FirstAlbumCheck(minFirst, maxFirst, artist) && MembersCheck(membersNum, artist) {
			if location != "" {
				if LocationCheck(location, artist) {
					NewData = append(NewData, Artists[i])
				}
			} else {
				NewData = append(NewData, Artists[i])
			}
		}
	}
	FiltredArtists = NewData
}

func IsUnique(id int, Data []Artist) bool {
	for _, m := range Data {
		if m.Id == id {
			return false
		}
	}
	return true
}

func CreationDateCheck(minCreate, maxCreate string, artist Artist) bool {
	return Atoi(minCreate) <= artist.CreationDate && Atoi(maxCreate) >= artist.CreationDate
}

func FirstAlbumCheck(minFirst, maxFirst string, artist Artist) bool {
	date := Atoi(artist.FirstAlbum[6:])
	return Atoi(minFirst) <= date && Atoi(maxFirst) >= date
}

func LocationCheck(location string, artist Artist) bool {
	for city := range artist.DatesLocations {
		city2 := strings.ReplaceAll(strings.ReplaceAll(city, "_", " "), "-", " ")
		if strings.Contains(strings.ToLower(city), strings.ToLower(location)) || strings.Contains(strings.ToLower(city2), strings.ToLower(location)) {
			return true
		}
	}
	return false
}

func MembersCheck(membersNum []string, artist Artist) bool {
	if membersNum == nil {
		return true
	}
	for i := range membersNum {
		if membersNum[i] == strconv.Itoa(len(artist.Members)) {
			return true
		}
	}
	return false
}

func Atoi(snum string) int {
	num, err := strconv.Atoi(snum)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return num
}
