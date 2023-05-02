package osrmclient

import (
	"context"
	geo "github.com/paulmach/go.geo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetDistance(t *testing.T) {

	src := *geo.NewPoint(13.388860, 52.517037)
	dst := []geo.Point{*geo.NewPoint(13.428555, 52.523219), *geo.NewPoint(13.397634, 52.529407)}

	routes, err := GetDistance(context.Background(), src, dst)
	require.NoError(t, err)

	require.Equal(t, float32(3804.3), routes[0].Distance)
	require.Equal(t, float32(389.1), routes[0].Duration)

	require.Equal(t, float32(1885.8), routes[1].Distance)
	require.Equal(t, float32(260.2), routes[1].Duration)
}
