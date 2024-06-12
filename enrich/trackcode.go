package enrich

//go:generate stringer -type TrackCode

// TrackCode is an enumerated type.
type TrackCode uint

const (
	UpMainFast                             TrackCode = 1100
	UpSlow                                 TrackCode = 1200
	UpGoods                                TrackCode = 1300
	UpSingle                               TrackCode = 1400
	UpLoop                                 TrackCode = 1500
	UpTerminal                             TrackCode = 1600
	UpCrossover                            TrackCode = 1700
	UpOtherOrEngine                        TrackCode = 1800
	UpSiding                               TrackCode = 1900
	DownMainFast                           TrackCode = 2100
	DownSlow                               TrackCode = 2200
	DownGoods                              TrackCode = 2300
	DownSingle                             TrackCode = 2400
	DownLoop                               TrackCode = 2500
	DownTerminal                           TrackCode = 2600
	DownCrossover                          TrackCode = 2700
	DownOtherOrEngine                      TrackCode = 2800
	DownSiding                             TrackCode = 2900
	ReversibleOrBiDirectionalMainFast      TrackCode = 3100
	ReversibleOrBiDirectionalSlow          TrackCode = 3200
	ReversibleOrBiDirectionalGoods         TrackCode = 3300
	ReversibleOrBidirectionalSingle        TrackCode = 3400
	ReversibleOrBiDirectionalLoop          TrackCode = 3500
	ReversibleOrBiDirectionalTerminal      TrackCode = 3600
	ReversibleOrBiDirectionalCrossover     TrackCode = 3700
	ReversibleOrBiDirectionalOtherOrEngine TrackCode = 3800
	ReversibleOrBiDirectionalSiding        TrackCode = 3900
)
