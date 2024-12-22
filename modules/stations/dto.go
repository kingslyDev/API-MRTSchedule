package stations

type Stations struct {
	Id string `json:"nid"`
	Name string `json:"title"`
}

type StationsResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
}