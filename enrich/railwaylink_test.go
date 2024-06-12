package enrich_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/enrich"
	"github.com/twpayne/go-shapefile"
)

func TestImportLinks(t *testing.T) {
	f, err := shapefile.Read("E:\\RailData\\NetworkRail\\ENRICH\\2024-01-16\\RailwayLinks", nil)
	assert.NoError(t, err)

	railwayLinks := make([]enrich.RailwayLink, f.NumRecords())

	for i := range f.NumRecords() {
		fields, geom := f.Record(i)

		railwayLinks[i], err = enrich.ParseRailwayLink(fields, geom)
		assert.NoError(t, err)
	}

	assert.Equal(t, f.NumRecords(), len(railwayLinks))
}
