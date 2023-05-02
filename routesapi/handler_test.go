package routesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := gin.Default()

	r.GET("/routes", GetFastestRoutes)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://your-service/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)

	expectedResponse := `{"Routes":[{"Destination":[13.428555,52.523219],"Duration":389.1,"Distance":3804.3},{"Destination":[13.397634,52.529407],"Duration":260.2,"Distance":1885.8}]}`
	require.Equal(t, expectedResponse, w.Body.String())
}
