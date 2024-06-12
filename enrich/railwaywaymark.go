package enrich

import (
	"errors"
	"time"

	"github.com/twpayne/go-geom"
)

type RailwayWaymark struct {
	// Unique identifier generated when the shapefile is created.
	FID string

	// Point – Represents the approximate location of a milepost or is a point added to support business needs.
	Shape geom.T

	// Unique ID which is consistent and maintained as far as possible.
	// ASSET_ID’s are retired/introduced when track layouts change significantly
	AssetID string

	// Engineers Line Reference, a railway term used to refer to a section of track covering multiple lines.
	ELR string

	// This denotes the master unit of measurement for the ELR:
	// K – Kilometres
	// M – Miles
	Unit string

	// This is the locational value in either miles or kilometres of point
	Value float64

	// This is the version number linked to the ASSET_ID which is increased each time
	// the data is amended in the main system, this is either its attribution or geometry.
	Version float64

	// The date the ‘RailwayWaymarks’ and corresponding information was extracted from the master data to create the
	// Open Data
	Extracted time.Time
}

func ParseRailwayWaymark(fields map[string]any, geom geom.T) (RailwayWaymark, error) {
	assetID, ok := fields["ASSET_ID"].(string)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse ASSET_ID")
	}

	elr, ok := fields["ASSET_ID"].(string)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse ELR")
	}

	unit, ok := fields["UNIT"].(string)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse UNIT")
	}

	value, ok := fields["VALUE"].(float64)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse VALUE")
	}

	version, ok := fields["VERSION"].(float64)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse VERSION")
	}

	extracted, ok := fields["EXTRACTED"].(time.Time)
	if !ok {
		return RailwayWaymark{}, errors.New("unable to parse EXTRACTED")
	}

	return RailwayWaymark{
		FID:       "",
		Shape:     geom,
		AssetID:   assetID,
		ELR:       elr,
		Unit:      unit,
		Value:     value,
		Version:   version,
		Extracted: extracted,
	}, nil
}
