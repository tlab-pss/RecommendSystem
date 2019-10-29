package location

import (
	"encoding/json"
	"net/http"
)

func GetAll() (*[]Location, error) {
	res, _ := http.Get("pd/api/locations")
	defer res.Body.Close()

	l := make([]Location, 0)
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		return nil, err
	}

	return &l, nil
}
