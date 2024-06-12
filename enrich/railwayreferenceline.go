package enrich

import (
	"errors"
	"time"

	"github.com/twpayne/go-geom"
)

type RailwayReferenceLine struct {
	// Unique identifier generated when the shapefile is created.
	FID string

	// Polyline – Represent the general location of the ELR.
	Shape geom.T

	// Unique ID which is consistent and maintained as far as possible. Each ELR has only one ASSET_ID.
	AssetID string

	// Engineers Line Reference, a railway term used to refer to a section of track covering multiple lines.
	ELR string

	// The start calibrated mileage value for the ELR The value is always in mile.yards irrespective
	// of the unit of measurement within the corresponding RailwayWaymark.shp data.
	Start float64

	// The end calibrated mileage value for the ELR. The value is always in mile.yards irrespective
	// of the unit of measurement within the corresponding RailwayWaymark.shp data.
	End float64

	// This is the version number linked to the ASSET_ID which is increased each time the data is
	// amended in the main system, this is either its attribution or geometry.
	Version float64

	// The date the ‘RailwayReferenceLines’ were extracted from the master data to create the
	// Open Data
	Extracted time.Time
}

func ParseRailwayReferenceLine(fields map[string]any, geom geom.T) (RailwayReferenceLine, error) {
	assetID, ok := fields["ASSET_ID"].(string)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse ASSET_ID")
	}

	elr, ok := fields["ELR"].(string)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse ELR")
	}

	start, ok := fields["START"].(float64)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse START")
	}

	end, ok := fields["END"].(float64)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse END")
	}

	version, ok := fields["VERSION"].(float64)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse VERSION")
	}

	extracted, ok := fields["EXTRACTED"].(time.Time)
	if !ok {
		return RailwayReferenceLine{}, errors.New("unable to parse EXTRACTED")
	}

	return RailwayReferenceLine{
		FID:       "",
		Shape:     geom,
		AssetID:   assetID,
		ELR:       elr,
		Start:     start,
		End:       end,
		Version:   version,
		Extracted: extracted,
	}, nil
}
