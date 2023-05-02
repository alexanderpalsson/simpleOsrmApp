package routesapi

import (
	osrm "github.com/gojuno/go.osrm"
	geo "github.com/paulmach/go.geo"
)

type route struct {
	Destination geo.Point
	Duration    float32
	Distance    float32
}

// routes contains a list of route and implements the go inbuilt sort interface which guarantees n*log(n) sorting
type routes struct {
	Routes []route
}

func (r routes) Len() int {
	return len(r.Routes)
}

func (r routes) Less(i, j int) bool {
	return r.Routes[i].Duration > r.Routes[j].Duration
}

func (r routes) Swap(i, j int) {
	r.Routes[i], r.Routes[j] = r.Routes[j], r.Routes[i]
}

func osmrRoutesToRoutes(dst []geo.Point, osmrRoutes []osrm.Route) routes {
	var rts routes
	for i, r := range osmrRoutes {
		rts.Routes = append(rts.Routes, route{
			Destination: dst[i],
			Distance:    r.Distance,
			Duration:    r.Duration,
		})
	}

	return rts
}
