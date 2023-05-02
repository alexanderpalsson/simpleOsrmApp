package routesapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ingrid/internal/osrmclient"
	geo "github.com/paulmach/go.geo"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func GetFastestRoutes(c *gin.Context) {
	src, err := parseSrc(c)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	dst, err := parseDst(c)
	if err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	osrmRoutes, err := osrmclient.GetDistance(context.Background(), src, dst)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	rspRoutes := osmrRoutesToRoutes(dst, osrmRoutes)
	sort.Sort(rspRoutes)

	c.JSON(http.StatusOK, rspRoutes)
}

func parseSrc(c *gin.Context) (geo.Point, error) {
	srcQuery, srcExist := c.GetQuery("src")
	if !srcExist {
		return geo.Point{}, errors.New("missing src in routesapi query")

	}

	src, err := parseFloats(strings.Split(srcQuery, ","))
	if err != nil {
		return geo.Point{}, errors.New("incorrect format of src, float format expected")
	}

	if len(src) != 2 {
		return geo.Point{}, errors.New("invalid src length, expected two points separated by a comma")
	}

	return *geo.NewPoint(src[0], src[1]), err
}

func parseDst(c *gin.Context) ([]geo.Point, error) {
	dstQuery, dstExist := c.GetQueryArray("dst")
	if !dstExist {
		return []geo.Point{}, errors.New("missing dst in routesapi query")
	}

	var dstPoints []geo.Point
	for _, dst := range dstQuery {
		dstFloat, err := parseFloats(strings.Split(dst, ","))
		if err != nil {
			return []geo.Point{}, errors.New("incorrect format of dst, float format expected")
		}

		if len(dstFloat) != 2 {
			return []geo.Point{}, errors.New("invalid dst length, expected even number of points separated by a commas")
		}

		dstPoints = append(dstPoints, *geo.NewPoint(dstFloat[0], dstFloat[1]))
	}

	return dstPoints, nil
}

func parseFloats(s []string) ([]float64, error) {
	var floats []float64
	for _, str := range s {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}

		floats = append(floats, f)
	}

	return floats, nil
}

func handleError(c *gin.Context, statusCode int, err error) {
	log.Println(fmt.Errorf("error handling routesapi call %w", c.AbortWithError(statusCode, err)))
}
