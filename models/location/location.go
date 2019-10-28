package location

type Location struct {
	ID        string  `json:"ID"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	CreatedAt string  `json:"CreatedAt"`
}
