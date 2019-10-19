package interactors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func SpecifyLocations() error {
	ls, err := getLocations()

	if err != nil {
		return err
	}

	home, office, err := specify(*ls)

	if err != nil {
		return err
	}

	bByte, _ := json.Marshal(map[string]interface{}{
		"house": map[string]interface{}{
			"lat": home.lat,
			"lnt": home.lng,
		},
		"office": map[string]interface{}{
			"lat": office.lat,
			"lnt": office.lng,
		},
	})

	req, err := http.NewRequest("POST", "pd/api/locations", bytes.NewReader(bByte))

	if err != nil {
		return err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func specify(ls []location) (pair, pair, error) {
	ht := make([]location, 0)
	ot := make([]location, 0)

	for _, l := range ls {
		t, err := time.Parse(time.RFC3339Nano, l.CreatedAt)

		if err != nil {
			return pair{}, pair{}, err
		}

		h := t.Hour() + 9

		// todo: 時間決め打ちしてるけど、将来的には人によってここも変えたい
		if 0 < h && h < 5 {
			ht = append(ht, l)
		}

		if 13 < h && h < 18 {
			ot = append(ot, l)
		}
	}

	//h := pair{}
	//for _, t := range ht {
	// todo: 最多要素を抽出したい
	//}
	h := pair{
		lat: ht[0].Latitude,
		lng: ht[0].Longitude,
	}

	//o := pair{}
	//for _, t := range ot {
	// todo: 最多要素を抽出したい
	//}
	o := pair{
		lat: ot[0].Latitude,
		lng: ot[0].Longitude,
	}

	return h, o, nil
}

func getLocations() (*[]location, error) {
	res, _ := http.Get("pd/api/locations")
	defer res.Body.Close()

	l := make([]location, 0)
	if err := json.NewDecoder(res.Body).Decode(&l); err != nil {
		return nil, err
	}

	return &l, nil
}


type pair struct {
	lat float64
	lng float64
}

type location struct {
	ID        string  `json:"ID"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	CreatedAt string  `json:"CreatedAt"`
}
