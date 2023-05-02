package osrmclient

import (
	"context"
	osrm "github.com/gojuno/go.osrm"
	geo "github.com/paulmach/go.geo"
)

const osrmURL = "https://router.project-osrm.org"

func GetShortestDistance(ctx context.Context, src geo.Point, dst []geo.Point) ([]osrm.Route, error) {
	client := osrm.NewFromURL(osrmURL)
	var routes []osrm.Route
	for _, d := range dst {
		resp, err := client.Route(ctx, osrm.RouteRequest{
			Profile: "car",
			Coordinates: osrm.NewGeometryFromPointSet(geo.PointSet{
				src,
				d,
			}),
			Geometries: osrm.GeometriesGeojson,
		})
		if err != nil {
			return []osrm.Route{}, err
		}

		routes = append(routes, resp.Routes[0])
	}

	return routes, nil
}
