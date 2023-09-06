package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Unmarshal() error {
	// if len(Artists) != 0 {
	// 	return nil
	// }
	Artistinfo, err := http.Get(ArtistApi)
	if err != nil {
		return err
	}

	defer Artistinfo.Body.Close()

	body, err := ioutil.ReadAll(Artistinfo.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Artists)
	if err != nil {
		fmt.Println("Can not unmarshall JSON")
		return err
	}

	return nil
}

func UnmarshalRelation() error {
	Relationsinfo, err := http.Get(ArtistDate)
	if err != nil {
		return err
	}

	defer Relationsinfo.Body.Close()

	body, err := ioutil.ReadAll(Relationsinfo.Body)
	Err := json.Unmarshal(body, &Relations)
	if Err != nil {
		return Err
	}
	for k := range Artists {
		Artists[k].DatesLocations = Relations.Index[k].DatesLocations
	}
	return nil
}
