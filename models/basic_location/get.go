package basic_location

import (
	"encoding/json"
	"net/http"
)

func Get() (*BasicLocation, error) {
	res, _ := http.Get("pd/api/basic-locations")
	defer res.Body.Close()

	bl := BasicLocation{}
	if err := json.NewDecoder(res.Body).Decode(&bl); err != nil {
		return nil, err
	}

	return &bl, nil
}
