package stations

type Service interface {
		GetAllStations() (response []StationsResponse, err error)
}
