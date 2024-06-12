package enrich_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/enrich"
	"github.com/twpayne/go-shapefile"
)

func TestImportRailwayMarks(t *testing.T) {
	f, err := shapefile.Read("E:\\RailData\\NetworkRail\\ENRICH\\2024-01-16\\RailwayWaymarks", nil)
	assert.NoError(t, err)

	waymarks := make([]enrich.RailwayWaymark, f.NumRecords())

	for i := range f.NumRecords() {
		fields, geom := f.Record(i)

		waymarks[i], err = enrich.ParseRailwayWaymark(fields, geom)
		assert.NoError(t, err)
	}

	assert.Equal(t, f.NumRecords(), len(waymarks))
}
