package enrich

import (
	"errors"
	"time"

	"github.com/twpayne/go-geom"
)

type RailwayLink struct {
	// Unique identifier generated when the shapefile is created.
	FID string

	// Polyline – Represents the approximate centreline of the railway track.
	Shape geom.T

	// Unique ID which is consistent and maintained as far as possible. ASSET_ID’s are
	// retired/introduced when track layouts change significantly.
	AssetID string

	// Engineers Line Reference, a railway term used to refer to a section of track covering
	// multiple lines.
	ELR string

	// A four digit code used within the railway to identify each section of track which indicates
	// speed, direction and type of use.
	TrackID string

	// The start calibrated mileage value for the link. The value is always in mile.yards irrespective
	// of the unit of measurement within the corresponding RailwayWaymark.shp data.
	Start float64

	// The end calibrated mileage value for the link. The value is always in mile.yards irrespective
	// of the unit of measurement within the corresponding RailwayWaymark.shp data.
	End float64

	// This is the version number linked to the ASSET_ID which is increased each time the data is
	// amended in the main system, this is either its attribution or geometry.
	Version float64

	// A concatenated comma separated data field which shows the percentage of track for each
	// link by distance, for which the owner is identified as: 'NETWORK RAIL', 'THIRD PARTY' or
	// 'UNCLASSIFIED'. ‘UNCLASSIFIED’ mostly refers to third party data where all information
	// is not available or where data is being updated and is awaiting confirmation from the
	// business. The percentage is to the nearest whole number and is rounded up to 1% to
	// ensure that the existence of different information is represented. For this reason, the
	// percentage values should only be used as an approximation and not as a definitive
	// convertible distance length.
	Owner string

	// The date the ‘RailwayLinks’ and corresponding information was extracted from the master
	// data to create the Open Data
	Extracted time.Time
}

func ParseRailwayLink(fields map[string]any, geom geom.T) (RailwayLink, error) {
	assetID, ok := fields["ASSET_ID"].(string)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse ASSET_ID")
	}

	elr, ok := fields["ELR"].(string)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse ELR")
	}

	end, ok := fields["END"].(float64)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse END")
	}

	extracted, ok := fields["EXTRACTED"].(time.Time)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse EXTRACTED")
	}

	owner, ok := fields["OWNER"].(string)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse OWNER")
	}

	start, ok := fields["START"].(float64)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse START")
	}

	trackID, ok := fields["TRACK_ID"].(string)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse TRACKID")
	}

	version, ok := fields["VERSION"].(float64)
	if !ok {
		return RailwayLink{}, errors.New("unable to parse ELR")
	}

	return RailwayLink{
		FID:       "",
		Shape:     geom,
		AssetID:   assetID,
		ELR:       elr,
		TrackID:   trackID,
		Start:     start,
		End:       end,
		Version:   version,
		Owner:     owner,
		Extracted: extracted,
	}, nil
}
