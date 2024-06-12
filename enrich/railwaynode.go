package enrich

import (
	"errors"
	"time"

	"github.com/twpayne/go-geom"
)

type RailwayNode struct {
	// Unique identifier generated when the shapefile is created.
	FID string

	// Point – Represent the start and end of every ‘RailwayLink’.
	Shape geom.T

	// Unique ID which is consistent and maintained as far as possible. ASSET_ID’s are
	// retired/introduced when track layouts change significantly.
	AssetID string

	// The number of links that connect to the point, this value can be between 1 and 4:
	// 1 – Tracks start/end or no connecting data held
	// 2 – Pseudo Node – no change in direction is not possible, attributional change only
	// 3 – Directional change possible
	// 4 – Centre of diamond crossing no directional change possible at this location although
	// attributional changes may occur.
	Valency float64

	// This is the version number linked to the ASSET_ID which is increased each time the data is
	// amended in the main system, either its attribution or geometry.
	Version float64

	// The date the ‘RailwayNodes’ were extracted from the master data to create the Open Data
	Extracted time.Time
}

func ParseRailwayNode(fields map[string]any, geom geom.T) (RailwayNode, error) {
	assetID, ok := fields["ASSET_ID"].(string)
	if !ok {
		return RailwayNode{}, errors.New("unable to parse ASSET_ID")
	}

	valency, ok := fields["VALENCY"].(float64)
	if !ok {
		return RailwayNode{}, errors.New("unable to parse VALENCY")
	}

	version, ok := fields["VERSION"].(float64)
	if !ok {
		return RailwayNode{}, errors.New("unable to parse VERSION")
	}

	extracted, ok := fields["EXTRACTED"].(time.Time)
	if !ok {
		return RailwayNode{}, errors.New("unable to parse EXTRACTED")
	}

	return RailwayNode{
		FID:       "",
		Shape:     geom,
		AssetID:   assetID,
		Valency:   valency,
		Version:   version,
		Extracted: extracted,
	}, nil
}
