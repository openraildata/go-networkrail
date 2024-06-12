package smart

//go:generate go run github.com/rickb777/enumeration/v3 -type StepType

// StepType is an enumerated type.
type StepType int

const (
	// Between is a move between directly adjacent berths, and is the preferred type
	// of movement. The time reported to TRUST is the time that the train enters the
	// 'to' berth.
	Between StepType = iota // json:"B"
	// From is used to record a time for a train going in either direction (up or
	// down) from the specified berth to any other berth. The time reported to TRUST
	// is the time that the train leaves the 'from' berth.
	From // json:"F"
	// To is the opposite of the 'F' step type. It is used to record a time for a
	// train from any berth to the specified berth. The time reported to TRUST is the
	// time that the train enters the 'to' berth.
	To // json:"T"
	// IntermediateFirst is used to specify the route a train is taking, usually when
	// departing a station or junction. For example, if a 'D' move is specified as
	// 0101 to 0407, and a train moves between berths 0101, 0203, 0305 and 0407, the
	// move will be reported for the time the train left the first berth.
	IntermediateFirst // json:"D"
	// Clearout is used to report on a movement where the only indication is a cancel
	// message. For example, when a train leaves Network Rail infrastructure and
	// moves in to a siding or area not covered by a train describer, the only
	// message that will be received is a Clearout (CB) message.
	Clearout // json:"C"
	// Interpose is the opposite of the 'C' step type. It is used to report on a
	// movement where the only indication is an interpose message. For example, when
	// a train arrives on Network Rail infrastructure from a siding. The time
	// reported to TRUST is the time that the interpose happened.
	Interpose // json:"I"
	// Intermediate is similar to the 'D' move, but usually used for arrivals. The
	// time reported to TRUST is the time the last berth step was made.
	Intermediate // json:"E"
)
