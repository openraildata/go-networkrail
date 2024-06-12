package enrich_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/enrich"
	"github.com/twpayne/go-shapefile"
)

func TestImportReferenceLines(t *testing.T) {
	f, err := shapefile.Read("E:\\RailData\\NetworkRail\\ENRICH\\2024-01-16\\RailwayReferenceLines", nil)
	assert.NoError(t, err)

	referenceLines := make([]enrich.RailwayReferenceLine, f.NumRecords())

	for i := range f.NumRecords() {
		fields, geom := f.Record(i)

		referenceLines[i], err = enrich.ParseRailwayReferenceLine(fields, geom)
		assert.NoError(t, err)
	}

	assert.Equal(t, f.NumRecords(), len(referenceLines))
}
