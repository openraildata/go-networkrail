package smart

//go:generate go run github.com/rickb777/enumeration/v3 -type Event

// Event is an enumerated type.
type Event uint

const (
	ArriveUp   Event = iota // json:"A"
	DepartUp                // json:"B"
	ArriveDown              // json:"C"
	DepartDown              // json:"D"
)
