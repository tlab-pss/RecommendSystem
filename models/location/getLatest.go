package location

import (
	"encoding/json"
	"net/http"
)

func GetLatest() (*Location, error){
	res, _ := http.Get("pd/api/locations/latest")
	defer res.Body.Close()

	l := Location{}
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		return nil, err
	}

	return &l, nil
}
