package enrich_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/openraildata/go-networkrail/enrich"
	"github.com/twpayne/go-shapefile"
)

func TestImportNodes(t *testing.T) {
	f, err := shapefile.Read("E:\\RailData\\NetworkRail\\ENRICH\\2024-01-16\\RailwayNodes", nil)
	assert.NoError(t, err)

	nodes := make([]enrich.RailwayNode, f.NumRecords())

	for i := range f.NumRecords() {
		fields, geom := f.Record(i)

		nodes[i], err = enrich.ParseRailwayNode(fields, geom)
		assert.NoError(t, err)
	}

	assert.Equal(t, f.NumRecords(), len(nodes))
}
